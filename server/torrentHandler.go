package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"

	"github.com/ppkavinda/drive-torrent/appError"
	"github.com/ppkavinda/drive-torrent/db"
)

func (s *Server) newURLHandler(w http.ResponseWriter, r *http.Request) {
	type Url struct {
		Url string
	}

	decoder := json.NewDecoder(r.Body)
	var url Url
	err := decoder.Decode(&url)
	if err != nil {
		fmt.Printf("newMagnetHandler: %+v\n", err)
		return
	}

	remote, err := http.Get(url.Url)
	if err != nil {
		fmt.Printf("Invalud remote url : %+v\n", err)
		appError.WriteAsJSON(w, err, "Invalid remote url", 0, 422)
		return
	}
	fileData, err := ioutil.ReadAll(remote.Body)
	if err != nil {
		fmt.Printf("Failed To download remote torrent: %+v\n", err)
		appError.WriteAsJSON(w, err, "Faild to download remote torrent", 0, 422)
		return
	}

	reader := bytes.NewBuffer(fileData)
	info, err := metainfo.Load(reader)
	if err != nil {
		fmt.Printf("Unable to read metaInfo : %+v\n", err)
		appError.WriteAsJSON(w, err, "Invalid metainfo of torrent file", 0, 422)
		return
	}

	spec := torrent.TorrentSpecFromMetaInfo(info)
	if err := s.engine.NewTorrentFromSpec(spec, GetUser(r).Email); err != nil {
		fmt.Printf("Torrent Error: %+v\n", err)
		appError.WriteAsJSON(w, err, "Unable to load torrent", 0, 422)
		return
	}
	// return nil
}

func (s *Server) newMagnetHandler(w http.ResponseWriter, r *http.Request) {
	type Magnet struct {
		Magnet string
	}

	decoder := json.NewDecoder(r.Body)
	var magnet Magnet
	err := decoder.Decode(&magnet)
	if err != nil {
		fmt.Printf("newMagnetHandler: %+v\n", err)
		return
	}

	err = s.engine.NewMagnet(magnet.Magnet, GetUser(r).Email)
	// fmt.Fprintf(w, "INFO4 %v", err)
	if err != nil {
		appError.WriteAsJSON(w, err, err.Error(), 0, 422)
	}
}

func newTorrentFormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/newTorrent.html")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}
	tmpl.Execute(w, GetUser(r))
}

func (s *Server) getTorrentsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("INFO %v\n", s.engine.GetTorrents())
}

func (s *Server) startTorrent(w http.ResponseWriter, r *http.Request) {
	type Hash struct {
		Hash string
	}
	decoder := json.NewDecoder(r.Body)
	var hash Hash
	err := decoder.Decode(&hash)
	if err != nil {
		fmt.Printf("startTorrent: %+v\n", err)
		return
	}
	err = s.engine.StartTorrent(hash.Hash)
	if err != nil {
		fmt.Printf("startTorrent2: %+v\n", err)
		return
	}
}
func (s *Server) stopTorrent(w http.ResponseWriter, r *http.Request) {
	type Hash struct {
		Hash string
	}
	decoder := json.NewDecoder(r.Body)
	var hash Hash
	err := decoder.Decode(&hash)
	if err != nil {
		fmt.Printf("stopTorrent: %+v\n", err)
		return
	}
	err = s.engine.Stop(hash.Hash)
	if err != nil {
		fmt.Printf("stopTorrent2: %+v\n", err)
		return
	}
}

func (s *Server) removeTorrent(w http.ResponseWriter, r *http.Request) {
	type Hash struct {
		Hash string
	}
	decoder := json.NewDecoder(r.Body)
	var hash Hash
	err := decoder.Decode(&hash)
	if err != nil {
		fmt.Printf("removeTorrentHandler: %+v\n", err)
		return
	}

	err = s.engine.Delete(hash.Hash)
	if err != nil {
		fmt.Printf("removeTorrentHandler: %+v\n", err)
		return
	}

	err = db.DeleteTorrent(hash.Hash, GetUser(r).Email)
	if err != nil {
		fmt.Printf("Delete Torrent: %+v\n", err)
		return
	}
	log.Printf("deleted torrent: %+v\n", hash.Hash)

}
