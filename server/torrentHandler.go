package server

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/ppkavinda/drive-torrent/profile"
)

func (s *Server) newURLHandler(w http.ResponseWriter, r *http.Request) {
	// r.Use(isLoggedIn)

	url := r.FormValue("url")
	remote, err := http.Get(url)
	if err != nil {
		fmt.Printf("Invalud remote url : %+v\n", err)
		// return nil
	}
	fileData, err := ioutil.ReadAll(remote.Body)
	if err != nil {
		fmt.Printf("Failed To download remote torrent: %+v\n", err)
		// return nil
	}

	reader := bytes.NewBuffer(fileData)
	info, err := metainfo.Load(reader)
	if err != nil {
		fmt.Printf("Unable to read metaInfo : %+v\n", err)
		// return nil
	}

	spec := torrent.TorrentSpecFromMetaInfo(info)
	if err := s.engine.NewTorrentFromSpec(spec); err != nil {
		fmt.Printf("Torrent Error: %+v\n", err)
		// return nil
	}
	// return nil
}

func (s *Server) newMagnetHandler(w http.ResponseWriter, r *http.Request) {
	magnet := r.FormValue("magnet")

	s.engine.NewMagnet(magnet, GetUser().Email)
	// fmt.Fprintf(w, "INFO4 %v", magnet)
}

func newTorrentFormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/newTorrent.html")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}
	tmpl.Execute(w, profile.User)
}

func (s *Server) getTorrentsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("INFO %v\n", s.engine.GetTorrents())
}
