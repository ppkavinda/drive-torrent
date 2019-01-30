package main

import (
	"github.com/ppkavinda/drive-torrent/server"
)

func main() {

	s := server.Server{
		Port:       "3000",
		ConfigPath: "drive-torrent.json",
	}
	s.Start()
}
