package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

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
		Config engine.Config
	}
}

// StartServer will start the http server
func (s *Server) StartServer() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

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
