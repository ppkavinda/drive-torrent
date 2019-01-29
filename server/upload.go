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

	"github.com/ppkavinda/drive-torrent/engine"
	"google.golang.org/api/drive/v3"
)

func uploadToDrive(d *drive.Service, description string,
	parentId string, file *engine.File) (*drive.File, error) {

	filePath := filepath.Join("./downloads", strings.TrimSpace(file.Path))

	fileName := filepath.Base(file.Path)

	extension := filepath.Ext(filePath)
	mimeType := mime.TypeByExtension(extension)

	input, err := os.Open(filePath)
	fmt.Printf("DONE %+v\n", input)
	if err != nil {
		fmt.Printf("An error occurred FILE.OPEN: %v\n", err)
		return nil, err
	}

	// Grab file info
	inputInfo, err := input.Stat()
	if err != nil {
		return nil, err
	}

	fmt.Println("Start upload")
	f := &drive.File{Name: fileName, Parents: []string{parentId}, Description: description, MimeType: mimeType}
	getRate := MeasureTransferRate()

	// progress call back
	showProgress := func(current, total int64) {
		fmt.Printf("Uploaded at %s, %s/%s\r", getRate(current), Comma(current), Comma(total))
	}

	r, err := d.Files.Create(f).ResumableMedia(context.Background(), input, inputInfo.Size(), mimeType).ProgressUpdater(showProgress).Do()
	fmt.Printf("DONE3 %+v\n", f.Name)
	fmt.Printf("DONE4 %+v\n", *r)
	if err != nil {
		fmt.Printf("An error occurred: %v\nFILEPATH: %+v\n", err, input.Name())
		return nil, err
	}

	fmt.Printf("HIT %v\n", r)
	// Total bytes transferred
	bytes := r.Size
	// Print information about uploaded file
	fmt.Printf("Uploaded '%s' at %s, total %s\n", r.Name, getRate(bytes), FileSizeFormat(bytes, false))
	fmt.Printf("Upload Done. ID : %s\n", r.Id)
	// fmt.Printf("File : %+v\n", r)
	return r, nil
}

func getOrCreateDriveFolder(d *drive.Service, folderName string, parentId string) string {
	folderId := ""
	if folderName == "" {
		return ""
	}

	q := fmt.Sprintf("name=\"%s\" and mimeType=\"application/vnd.google-apps.folder\" and trashed = false ", folderName)
	if parentId != "" {
		q = fmt.Sprintf("name=\"%s\" and \"%s\" in parents and mimeType=\"application/vnd.google-apps.folder\" and trashed = false ", folderName, parentId)
	}

	r, err := d.Files.List().Q(q).PageSize(1).Do()
	if err != nil {
		fmt.Printf("%s\n", folderName)
		log.Fatalf("Unable to retrieve foldername. %+v", err)
	}

	if len(r.Files) > 0 {
		// fmt.Printf("%+v\n", r.Files[0])
		folderId = r.Files[0].Id
	} else {
		// no folder found create new
		fmt.Printf("Folder not found. Create new folder : %s\n", folderName)
		f := &drive.File{Name: folderName, MimeType: "application/vnd.google-apps.folder"}

		if parentId != "" {
			f = &drive.File{Name: folderName, Parents: []string{parentId}, MimeType: "application/vnd.google-apps.folder"}
		}
		r, err := d.Files.Create(f).Do()
		if err != nil {
			fmt.Printf("An error occurred when create folder: %v\n", err)
		}
		folderId = r.Id
	}
	return folderId
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
