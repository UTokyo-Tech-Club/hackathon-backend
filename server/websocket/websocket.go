package websocket

import (
	"encoding/json"
	"fmt"
	firebaseAuth "hackathon-backend/firebase"
	"hackathon-backend/utils/logger"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type ClientObject struct {
	Email           string `json:"email,omitempty"`
	Username        string `json:"userName,omitempty"`
	AuthToken       string `json:"authToken,omitempty"`
	ClientWebSocket *websocket.Conn
}

type WebSocketServer struct {
	Clients          map[*ClientObject]bool
	ClientTokenMap   map[string]*ClientObject
	registerClient   chan *ClientObject
	unregisterClient chan *ClientObject

	getUpgrader websocket.Upgrader

	lock sync.Mutex
}

func Init() *WebSocketServer {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return &WebSocketServer{
		Clients:          make(map[*ClientObject]bool),
		ClientTokenMap:   make(map[string]*ClientObject),
		registerClient:   make(chan *ClientObject),
		unregisterClient: make(chan *ClientObject),

		getUpgrader: upgrader,

		lock: sync.Mutex{},
	}
}

func (wss *WebSocketServer) SetupRoutes() {
	http.HandleFunc("/", wss.handleHomePage)
	http.HandleFunc("/ws", wss.handleEndPoint)
}

func (wss *WebSocketServer) SetupEventListeners() {

	// Client Registration
	go func() {
		for {
			select {
			case client := <-wss.registerClient:
				wss.lock.Lock()
				wss.Clients[client] = true
				wss.ClientTokenMap[client.AuthToken] = client
				wss.lock.Unlock()
			case client := <-wss.unregisterClient:
				wss.lock.Lock()
				if _, ok := wss.Clients[client]; ok {
					delete(wss.Clients, client)
					client.ClientWebSocket.Close()
				}
				wss.lock.Unlock()
			}
		}
	}()

}

func (wss *WebSocketServer) handleHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

// func closeConnection(wss *WebSocketServer, client *ClientObject) {
// 	logger := logger.NewLogger()
// 	err := client.ClientWebSocket.Close()
// 	if err != nil {
// 		logger.Error(err)
// 	}
// 	delete(wss.Clients, client)
// 	delete(wss.ClientTokenMap, client.AuthToken)
// }

func (wss *WebSocketServer) handleEndPoint(w http.ResponseWriter, r *http.Request) {
	logger.Info("WebSocket Endpoint Hit")

	wss.getUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := wss.getUpgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(err)
		return
	}

	fb := firebaseAuth.Init()
	isAuth := false

	client := &ClientObject{}
	defer client.ClientWebSocket.Close()

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			logger.Error(err)
			break
		}

		var msg *Message
		err = json.Unmarshal(p, &msg)
		if err != nil {
			logger.Error(err)
			break
		}

		// Guard until authentication
		if !isAuth {

			if msg.Type != "auth" {
				logger.Error("Must authenticate first")
				break
			}

			logger.Info("AuthToken: ", msg.Data)
			idToken, err := firebaseAuth.ValidateToken(fb, msg.Data)
			if err != nil {
				logger.Error("Error verifying token:", err)
				return
			}

			isAuth = true
			client.Username = idToken.UID
			wss.registerClient <- client

			logger.Info("Authenticated user: ", idToken.UID)
		}

	}

}
