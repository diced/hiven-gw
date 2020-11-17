package gateway

import (
	"sync"

	"github.com/gorilla/websocket"
)

// Websocket struct for ws & ws urls
type Websocket struct {
	URL    string
	Socket *websocket.Conn
	M      sync.Mutex
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
	w.M.Lock()
	defer w.M.Unlock()

	w.Socket.WriteJSON(Map{
		"op": code,
	})
}

// SendOPData sends an opcode with data to socket
func (w *Websocket) SendOPData(code int, data Map) {
	w.M.Lock()
	defer w.M.Unlock()

	w.Socket.WriteJSON(Map{
		"op": code,
		"d":  data,
	})
}

// Reconnect func
func (w *Websocket) Reconnect(token string) {
	w.SendOPData(2, Map{
		"token": token,
	})
}

// Heartbeat sends op code 3, used in a go routine
func (w *Websocket) Heartbeat() {
	w.SendOP(3)
}
