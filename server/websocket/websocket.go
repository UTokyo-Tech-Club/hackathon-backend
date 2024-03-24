package websocket

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type ClientObject struct {
	JoinedAt        time.Time `json:"joinedAt,omitempty"`
	IPAddress       string    `json:"ipAddress,omitempty"`
	Username        string    `json:"userName,omitempty"`
	EntryToken      string    `json:"entryToken,omitempty"`
	ClientWebSocket *websocket.Conn
}

type WebSocketServer struct {
	clients         map[*ClientObject]bool
	clientTokenMap  map[string]*ClientObject
	requestUpgrader websocket.Upgrader
}

func Init() *WebSocketServer {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return &WebSocketServer{
		clients:         make(map[*ClientObject]bool),
		clientTokenMap:  make(map[string]*ClientObject),
		requestUpgrader: upgrader,
	}
}

func (wss *WebSocketServer) SetupServer() {
	http.HandleFunc("/", wss.handleHomePage)
}

func (wss *WebSocketServer) handleHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}
