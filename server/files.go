package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

func (s *Server) uploadFiles(torrent *engine.Torrent) {
	email := getEmailOfTorrent(torrent)

	fmt.Printf("FINISHED: %s:%s\n", torrent.InfoHash, email)
	files := s.engine.GetFiles(torrent.InfoHash)
	client := getClient(OAuthConfig, email)

	srv, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	// var parentName []string
	parentName := strings.Split(files[0].Path, "/")
	if len(files) > 1 {

		for _, file := range files {
			// fullPath := filepath.Join("./downloads", file.Path)
			fileName := filepath.Base(file.Path)
			folders := strings.TrimSuffix(file.Path, "/"+fileName)

			parentID := getOrCreateDriveFolder(srv, "drive-torrent", "")
			for _, fldrName := range strings.Split(folders, "/") {
				parentID = getOrCreateDriveFolder(srv, fldrName, parentID)
			}
			_, err = uploadToDrive(srv, "", parentID, file)
			if err != nil {
				fmt.Printf("%+v\n", err)
			}
		}
	} else {
		parentID := getOrCreateDriveFolder(srv, "drive-torrent", "")
		_, err = uploadToDrive(srv, "", parentID, files[0])
		if err != nil {
			fmt.Printf("%+v\n", err)
		}
	}
	os.RemoveAll(filepath.Join("./downloads", parentName[0]))
}
