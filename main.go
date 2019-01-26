package main

import (
	"github.com/ppkavinda/drive-torrent/server"
)

func main() {
	// oauthConfig = configureOAuthClient("404364039745-0caba0fvhaja2cogru4jvl0gqq3anf50.apps.googleusercontent.com", "zRly0iH-ThMZrYRxER5PT_ue")

	// conf := torrent.Config{
	// 	DataDir: "./downloads",
	// 	DHTConfig: dht.ServerConfig{
	// 		StartingNodes: dht.GlobalBootstrapAddrs,
	// 	},
	// }
	// c, _ := torrent.NewClient(&conf)
	// defer c.Close()
	// t, _ := c.AddMagnet("magnet:?xt=urn:btih:187D2FA2CD25E7256BC2101B8CAC2EDAEC82994B&dn=Quick+Heal+Total+Security+key+%28Renewal+for+1+year%29+%5BRxV%5D&tr=http%3A%2F%2Ffr33dom.h33t.com%3A3310%2Fannounce&tr=http%3A%2F%2Fwww.cbtorrents.com%3A2710%2Fannounce&tr=http%3A%2F%2Fnemesis.1337x.org%2Fannounce&tr=http%3A%2F%2Fexodus.1337x.org%2Fannounce&tr=http%3A%2F%2Ftracker.thepiratebay.org%2Fannounce&tr=http%3A%2F%2Ftracker.thepiratebay.org%3A80%2Fannounce&tr=http%3A%2F%2Fgenesis.1337x.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.1337x.org%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.zer0day.to%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fcoppersurfer.tk%3A6969%2Fannounce")
	// x := <-t.GotInfo()
	// t.DownloadAll()
	// c.WaitAll()

	// fmt.Println(x)
	// fmt.Println(t.Stats())
	// fmt.Println(c.Torrents())
	s := server.Server{
		Port: "8080",
	}
	s.StartServer()
}
