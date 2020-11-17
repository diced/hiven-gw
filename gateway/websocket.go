package gateway

import (
	"log"

	"github.com/gorilla/websocket"
)

// Websocket struct for ws & ws urls
type Websocket struct {
	URL    string
	Socket *websocket.Conn
}

// NewWebsocket creates a new gateway websocket struct ^
func NewWebsocket(url string) Websocket {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		panic(err)
	}

	return Websocket{
		URL:    url,
		Socket: ws,
	}
}

// SendOP sends an opcode to socket
func (w *Websocket) SendOP(code int) {
	w.Socket.WriteJSON(map[string]interface{}{
		"op": code,
	})
}

// Reconnect func
func (w *Websocket) Reconnect(token string) {
	w.Socket.WriteJSON(map[string]interface{}{
		"op": 2,
		"d": map[string]interface{}{
			"token": token,
		},
	})
}

// Heartbeat sends op code 3, used in a go routine
func (w *Websocket) Heartbeat() {
	log.Println("hb")
	w.SendOP(3)
}