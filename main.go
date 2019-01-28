package main

import (
	"github.com/ppkavinda/drive-torrent/server"
)

func main() {

	s := server.Server{
		Port:       "8080",
		ConfigPath: "drive-torrent.json",
	}
	s.StartServer()
}
