package server

import (
	"time"

	"github.com/ppkavinda/drive-torrent/db"
	"github.com/ppkavinda/drive-torrent/engine"
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

func (s *Server) getTorrentsOfEmail(email string) []engine.Torrent {
	hashes := db.GetHashesOfEmail(email)
	var torrents []engine.Torrent
	torrents = make([]engine.Torrent, 0)

	for _, hash := range hashes {
		// torrents = append(torrents, s.state.Torrents[hash])
		torrent := s.state.Torrents[hash]
		if torrent != nil {
			torrents = append(torrents, *torrent)
		}
	}

	return torrents
}
