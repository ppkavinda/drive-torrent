package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// WsStruct : websocket data object
type WsStruct struct {
	Event string
	Data  interface{}
}

const (
	// Time allowed to write the file to the client.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Poll file for changes with this period.
	torrentPeriod = 2 * time.Second

	loginPeriod = 3 * time.Second
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// SocketHandler : handler for socket connection
func (s *Server) SocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			fmt.Printf("SocketHandler: %+v", err)
		}
		return
	}

	go writer(ws, r, s)
	reader(ws)
}

func writer(ws *websocket.Conn, r *http.Request, s *Server) {
	pingTicker := time.NewTicker(pingPeriod)
	torrentTicker := time.NewTicker(torrentPeriod)
	loginTicker := time.NewTicker(loginPeriod)
	fmt.Printf("ws Write\n")

	defer func() {
		pingTicker.Stop()
		torrentTicker.Stop()
		ws.Close()
	}()

	for {
		select {
		case <-torrentTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			email := GetUser(r).Email
			torrents := s.getTorrentsOfEmail(email)
			js, err := json.Marshal(WsStruct{
				"sync-torrent",
				torrents,
			})
			err = ws.WriteMessage(websocket.TextMessage, js)
			if err != nil {
				return
			}
		case <-loginTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			user := GetUser(r)
			js, err := json.Marshal(WsStruct{
				"sync-login",
				user,
			})
			err = ws.WriteMessage(websocket.TextMessage, js)
			if err != nil {
				return
			}

		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			err := ws.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				return
			}
		}
	}
}

func reader(ws *websocket.Conn) {
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}
