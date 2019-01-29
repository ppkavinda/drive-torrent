package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	// _ "github.com/mattn/go-sqlite3"
	"github.com/ppkavinda/drive-torrent/engine"
)

// Server stores config of server
type Server struct {
	Title      string
	Port       string
	engine     *engine.Engine
	Host       string
	ConfigPath string
	state      struct {
		sync.Mutex
		Config    engine.Config
		Downloads *fsNode
		Torrents  map[string]*engine.Torrent
	}
}

// StartServer will start the http server
func (s *Server) StartServer() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	setupDB()

	s.engine = engine.New()

	c := engine.Config{
		DownloadDirectory: "./downloads",
		EnableUpload:      true,
		AutoStart:         true,
	}

	if c.IncomingPort <= 0 || c.IncomingPort >= 65535 {
		c.IncomingPort = 50007
	}

	if err := s.reconfig(c); err != nil {
		return err.Error
		// return appErrorf(err, "Unable to Configure %v", err)
	}
	//poll torrents and files
	go func() {
		for {
			s.state.Lock()
			s.state.Torrents = s.engine.GetTorrents()
			// s.state.Downloads = s.listFiles()
			s.state.Unlock()
			time.Sleep(1 * time.Second)

			for _, torrent := range s.state.Torrents {
				if torrent.Finished && !torrent.Uploaded {
					s.uploadFile(torrent)
					torrent.Uploaded = true
				}
			}
		}
	}()

	host := s.Host
	if host == "" {
		host = "0.0.0.0"
	}

	registerRoutes(s)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	return nil
}

func (s *Server) reconfig(c engine.Config) *appError {
	dldir, err := filepath.Abs(c.DownloadDirectory)
	if err != nil {
		return appErrorf(err, "Invalid Path %v", err)
	}
	c.DownloadDirectory = dldir
	if err := s.engine.Config(c); err != nil {
		return appErrorf(err, "Unable to configure: %v", err)
	}
	b, _ := json.MarshalIndent(&c, "", " ")
	ioutil.WriteFile(s.ConfigPath, b, 0755)
	s.state.Config = c
	return nil
}

func setupDB() {

	db, err := sql.Open("sqlite3", "./info.db")
	if err != nil {
		fmt.Printf("SQL: %v", err)
	}
	defer db.Close()

	sqlStmt := `
		create table if not exists torrents (
			id integer auto_increment not null primary key,
			name text,
			hash text,
			email text
			);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	rows, err := db.Query("select id, name from torrents")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}

}

// http://blog.golang.org/error-handling-and-go
type appHandler func(http.ResponseWriter, *http.Request) *appError

type appError struct {
	Error   error
	Message string
	Code    int
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		log.Printf("Handler error: status code: %d, message: %s, underlying err: %#v",
			e.Code, e.Message, e.Error)

		http.Error(w, e.Message, e.Code)
	}
}

func appErrorf(err error, format string, v ...interface{}) *appError {
	return &appError{
		Error:   err,
		Message: fmt.Sprintf(format, v...),
		Code:    500,
	}
}
