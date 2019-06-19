package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetAllTorrents : end point for get all torrents (admin)
func (s *Server) getAllTorrents(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v", s.engine.Torrents())

	js, err := json.Marshal(s.engine.Torrents())
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	w.Header().Set("content-type", "application/json")
	w.Write(js)
}
