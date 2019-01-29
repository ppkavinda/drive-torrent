package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/ppkavinda/drive-torrent/engine"
	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
)

type fsNode struct {
	Name     string
	Size     int64
	Modified time.Time
	Children []*fsNode
}

// func (s *Server) listFiles() *fsNode {
// 	rootDir := s.state.Config.DownloadDirectory
// 	root := &fsNode{}

// 	if info, err := os.Stat(rootDir); err == nil {
// if err := list(rootDir, info, root, new(int)); err != nil {
// 	fmt.Printf("%v", err)
// }
// }
// return root
// }

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func getEmailOfTorrent(torrent *engine.Torrent) string {
	db, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("SQL: %v", err)
	}
	stmt, err := db.Prepare("select id, email from torrents where hash = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	defer db.Close()

	var id, email string
	err = stmt.QueryRow(torrent.InfoHash).Scan(&id, &email)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("delete from torrents where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("FINISHED : %s\n", email)
	return email
}

func (s *Server) uploadFile(torrent *engine.Torrent) {
	email := getEmailOfTorrent(torrent)

	fmt.Printf("FINISHED: %s:%s\n", torrent.InfoHash, email)
	files := s.engine.GetFiles(torrent.InfoHash)
	// var parentName []string
	if len(files) > 1 {
		// parentName = strings.Split(files[0].Path, "/")
	} else {
		client := getClient(OAuthConfig, email)

		srv, err := drive.New(client)
		if err != nil {
			log.Fatal("Unable to retrieve Drive client: %v", err)
		}
		inputFile := filepath.Join("./downloads", files[0].Path)
		outputTitle := filepath.Base(files[0].Path)

		extension := filepath.Ext(inputFile)
		mimeType := mime.TypeByExtension(extension)

		_, err = uploadToDrive(srv, outputTitle, "", "ParentFolderName", mimeType, inputFile)
		if err != nil {
			fmt.Printf("%+v", err)
		}
		os.Remove(inputFile)
	}

}

func uploadToDrive(d *drive.Service, title string, description string,
	parentName string, mimeType string, filename string) (*drive.File, error) {
	input, err := os.Open(filename)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return nil, err
	}

	// Grab file info
	inputInfo, err := input.Stat()
	if err != nil {
		return nil, err
	}

	// parentId := getOrCreateFolder(d, parentName)

	fmt.Println("Start upload")
	f := &drive.File{Name: title, Description: description, MimeType: mimeType}
	// if parentId != "" {
	// 	p := &drive.ParentReference{Id: parentId}
	// 	f.Parents = []*drive.ParentReference{p}
	// }
	getRate := MeasureTransferRate()

	// progress call back
	showProgress := func(current, total int64) {
		fmt.Printf("Uploaded at %s, %s/%s\r", getRate(current), Comma(current), Comma(total))
	}

	r, err := d.Files.Create(f).ResumableMedia(context.Background(), input, inputInfo.Size(), mimeType).ProgressUpdater(showProgress).Do()
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return nil, err
	}

	// Total bytes transferred
	bytes := r.Size
	// Print information about uploaded file
	fmt.Printf("Uploaded '%s' at %s, total %s\n", r.Name, getRate(bytes), FileSizeFormat(bytes, false))
	fmt.Printf("Upload Done. ID : %s\n", r.Id)
	return r, nil
}

func getOrCreateFolder(d *drive.Service, folderName string) string {
	folderId := ""
	if folderName == "" {
		return ""
	}
	q := fmt.Sprintf("title=\"%s\" and mimeType=\"application/vnd.google-apps.folder\"", folderName)

	r, err := d.Files.List().Q(q).PageSize(1).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve foldername.", err)
	}

	if len(r.Files) > 0 {
		folderId = r.Files[0].Id
	} else {
		// no folder found create new
		fmt.Printf("Folder not found. Create new folder : %s\n", folderName)
		f := &drive.File{Name: folderName, Description: "Auto Create by gdrive-upload", MimeType: "application/vnd.google-apps.folder"}
		r, err := d.Files.Create(f).Do()
		if err != nil {
			fmt.Printf("An error occurred when create folder: %v\n", err)
		}
		folderId = r.Id
	}
	return folderId
}

func MeasureTransferRate() func(int64) string {
	start := time.Now()

	return func(bytes int64) string {
		seconds := int64(time.Now().Sub(start).Seconds())
		if seconds < 1 {
			return fmt.Sprintf("%s/s", FileSizeFormat(bytes, false))
		}
		bps := bytes / seconds
		return fmt.Sprintf("%s/s", FileSizeFormat(bps, false))
	}
}
func FileSizeFormat(bytes int64, forceBytes bool) string {
	if forceBytes {
		return fmt.Sprintf("%v B", bytes)
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}

	var i int
	value := float64(bytes)

	for value > 1000 {
		value /= 1000
		i++
	}
	return fmt.Sprintf("%.1f %s", value, units[i])
}

func Comma(v int64) string {
	sign := ""
	if v < 0 {
		sign = "-"
		v = 0 - v
	}

	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for v > 999 {
		parts[j] = strconv.FormatInt(v%1000, 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		v = v / 1000
		j--
	}
	parts[j] = strconv.Itoa(int(v))
	return sign + strings.Join(parts[j:], ",")
}
