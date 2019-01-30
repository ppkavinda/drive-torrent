package server

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config, email string) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := filepath.Join("./tokens/", email+".json")
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		fmt.Printf("Couldn't get Token\n")
		return nil
		// tok = getTokenFromWeb(config)
		// saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

func getOrCreateDriveFolder(d *drive.Service, folderName string, parentID string) string {
	folderID := ""
	if folderName == "" {
		return ""
	}

	q := fmt.Sprintf("name=\"%s\" and mimeType=\"application/vnd.google-apps.folder\" and trashed = false ", folderName)
	if parentID != "" {
		q = fmt.Sprintf("name=\"%s\" and \"%s\" in parents and mimeType=\"application/vnd.google-apps.folder\" and trashed = false ", folderName, parentID)
	}

	r, err := d.Files.List().Q(q).PageSize(1).Do()
	if err != nil {
		fmt.Printf("%s\n", folderName)
		log.Fatalf("Unable to retrieve foldername. %+v", err)
	}

	if len(r.Files) > 0 {
		folderID = r.Files[0].Id
	} else {
		// no folder found create new
		fmt.Printf("Folder not found. Create new folder : %s\n", folderName)
		f := &drive.File{Name: folderName, MimeType: "application/vnd.google-apps.folder"}

		if parentID != "" {
			f = &drive.File{Name: folderName, Parents: []string{parentID}, MimeType: "application/vnd.google-apps.folder"}
		}
		r, err := d.Files.Create(f).Do()
		if err != nil {
			fmt.Printf("An error occurred when create folder: %v\n", err)
		}
		folderID = r.Id
	}
	return folderID
}
