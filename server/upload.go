package server

import (
	"context"
	"fmt"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/ppkavinda/drive-torrent/db"
	"github.com/ppkavinda/drive-torrent/engine"
	"google.golang.org/api/drive/v3"
)

func (s *Server) uploadFiles(infohash string) {
	emails := db.GetEmailOfTorrent(infohash)
	files := s.engine.GetFiles(infohash)
	parentName := strings.Split(files[0].Path, "/")
	torrent, err := s.engine.GetTorrent(infohash)
	if err != nil {
		fmt.Printf("Error: cannot get torrent of %s %+v\n", infohash, err)
	}

	for _, email := range emails {
		log.Printf("uploading to :: %+v\n", email)
		// continue

		client := getClient(OAuthConfig, email)

		srv, err := drive.New(client)
		if err != nil {
			log.Fatalf("Unable to retrieve Drive client: %v\n", err)
		}

		// var parentName []string
		if len(files) > 1 {

			for _, file := range files {
				// fullPath := filepath.Join("./downloads", file.Path)
				fileName := filepath.Base(file.Path)
				folders := strings.TrimSuffix(file.Path, "/"+fileName)

				parentID := getOrCreateDriveFolder(srv, "drive-torrent", "")
				for _, fldrName := range strings.Split(folders, "/") {
					parentID = getOrCreateDriveFolder(srv, fldrName, parentID)
				}
				_, err = uploadToDrive(srv, "", parentID, file, &torrent)
				if err != nil {
					fmt.Printf("%+v\n", err)
				}
			}
		} else {
			parentID := getOrCreateDriveFolder(srv, "drive-torrent", "")
			_, err = uploadToDrive(srv, "", parentID, files[0], &torrent)
			if err != nil {
				fmt.Printf("%+v\n", err)
			}
		}
		err = db.DeleteTorrent(infohash, email)
		if err != nil {
			fmt.Printf("%+v\n", err)
		}
	}
	err = os.RemoveAll(filepath.Join("./downloads", parentName[0]))
	if err != nil {
		fmt.Printf("Cannot Delete file %+v\n", err)
		return
	}

	err = s.engine.Delete(infohash)
	if err != nil {
		fmt.Printf("Cannot Delete file %+v\n", err)
		return
	}
}

func uploadToDrive(d *drive.Service, description string,
	parentID string, file *engine.File, torrent **engine.Torrent) (*drive.File, error) {

	filePath := filepath.Join("./downloads", strings.TrimSpace(file.Path))

	fileName := filepath.Base(file.Path)

	extension := filepath.Ext(filePath)
	mimeType := mime.TypeByExtension(extension)

	input, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("An error occurred FILE.OPEN: %v\n", err)
		return nil, err
	}

	// Grab file info
	inputInfo, err := input.Stat()
	if err != nil {
		return nil, err
	}

	log.Println("Start upload")
	f := &drive.File{Name: fileName, Parents: []string{parentID}, Description: description, MimeType: mimeType}
	getRate := MeasureTransferRate()

	// progress call back
	showProgress := func(current, total int64) {
		(*torrent).UploadRate = getRate(current)
		(*torrent).Uploaded = current
		// log.Printf("%+v", torrent)

		// log.Printf("Uploaded at %s, %s/%s\r", getRate(current), Comma(current), Comma(total))
	}

	r, err := d.Files.Create(f).ResumableMedia(context.Background(), input, inputInfo.Size(), mimeType).ProgressUpdater(showProgress).Do()
	if err != nil {
		fmt.Printf("An error occurred: %v\nFILEPATH: %+v\n", err, input.Name())
		return nil, err
	}

	// Total bytes transferred
	bytes := r.Size
	// Print information about uploaded file
	log.Printf("Uploaded '%s' at %s, total %s\n", r.Name, getRate(bytes), FileSizeFormat(bytes, false))
	// fmt.Printf("File : %+v\n", r)
	return r, nil
}

// MeasureTransferRate : returns the transfer rate
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

// FileSizeFormat : get the file size format MB, KB, etc
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

// Comma : add commas appropriatly
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
