package server

import (
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/net/websocket"
)

type WsStruct struct {
	Event string
	Data  interface{}
}

// SocketHandler : handler for socket connection
func (s *Server) SocketHandler(ws *websocket.Conn) {
	// var err error
	// syncLogin(s, ws)
	syncInitialTorrents(s, ws)
	// syncNewTorrent(s, ws)

	// for {
	// 	var reply string

	// 	if err = websocket.Message.Receive(ws, &reply); err != nil {
	// 		fmt.Println("Can't receive")
	// 		break
	// 	}

	// 	fmt.Println("Received back from client: " + reply)

	// 	msg := "Received:  "
	// 	fmt.Println("Sending to client: " + msg)

	// 	if err = websocket.Message.Send(ws, msg); err != nil {
	// 		fmt.Println("Can't send")
	// 		break
	// 	}

	// }
}

func syncNewTorrent(s *Server, ws *websocket.Conn) {

}

func syncInitialTorrents(s *Server, ws *websocket.Conn) {
	func() {
		for {
			email := GetUser().Email
			torrents := s.getTorrentsOfEmail(email)
			js, err := json.Marshal(WsStruct{
				"sync-torrent",
				torrents,
			})
			if err != nil {
				fmt.Printf("Can't send1 %+v\n", err)
				break
			}
			if err := websocket.Message.Send(ws, string(js)); err != nil {
				fmt.Printf("Can't send2 %+v\n", err)
				break
			}
			js, err = json.Marshal(WsStruct{
				"login-status",
				GetUser(),
			})
			if err = websocket.Message.Send(ws, string(js)); err != nil {
				fmt.Printf("Can't send2 %+v\n", err)
				break
			}

			// fmt.Printf("sendTorrents: %+v\n", torrents)
			// fmt.Printf("sendTorrents: js %s\n", js)
			fmt.Print("syncTorrent\n")
			time.Sleep(2 * time.Second)
		}
	}()
}

func syncLogin(s *Server, ws *websocket.Conn) {
	go func() {
		for {
			time.Sleep(2 * time.Second)

			one := []byte{}
			// // ws.SetReadDeadline(time.Now())
			if _, err := ws.Read(one); err != nil {
				fmt.Printf("detected closed LAN connection %+v\n", err)
				// 	// ws.Close()
				// 	// ws = nil
				break
			}

			fmt.Printf("syncLogin\n")
		}
	}()
}
