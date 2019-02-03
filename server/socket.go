package server

import (
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/net/websocket"
)

func (s *Server) SocketHandler(ws *websocket.Conn) {
	// var err error

	sendTorrents(s, ws)

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

func sendTorrents(s *Server, ws *websocket.Conn) {
	func() {
		email := GetUser().Email
		for {
			torrents := s.getTorrentsOfEmail(email)
			js, err := json.Marshal(torrents)
			if err != nil {
				fmt.Printf("Can't send1 %+v\n", err)
				break
			}
			if err := websocket.Message.Send(ws, string(js)); err != nil {
				fmt.Printf("Can't send2 %+v\n", err)
				break
			}
			fmt.Printf("sendTorrents: %+v\n", torrents)
			fmt.Printf("sendTorrents: js %s\n", js)
			time.Sleep(5 * time.Second)
		}
	}()
}
