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
	parentID string, file *engine.File) (*drive.File, error) {

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
		log.Printf("Uploaded at %s, %s/%s\r", getRate(current), Comma(current), Comma(total))
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
