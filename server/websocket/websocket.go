package websocket

import (
	"fmt"
	firebaseAuth "hackathon-backend/firebase"
	"hackathon-backend/server"
	"hackathon-backend/utils/logger"
	"net/http"
	"sync"

	"firebase.google.com/go/auth"
	"github.com/gorilla/websocket"
)

type ClientObject struct {
	UID             string `json:"uid"`
	ClientWebSocket *websocket.Conn
}

type WebSocketServer struct {
	Clients          map[*ClientObject]bool
	ClientUIDMap     map[string]*ClientObject
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
		ClientUIDMap:     make(map[string]*ClientObject),
		registerClient:   make(chan *ClientObject),
		unregisterClient: make(chan *ClientObject),

		getUpgrader: upgrader,

		lock: sync.Mutex{},
	}
}

func (wss *WebSocketServer) SetupRouter() {
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
				wss.ClientUIDMap[client.UID] = client
				wss.lock.Unlock()
			case client := <-wss.unregisterClient:
				wss.lock.Lock()
				if _, ok := wss.Clients[client]; ok {
					delete(wss.Clients, client)
					delete(wss.ClientUIDMap, client.UID)
					if client.ClientWebSocket != nil {
						client.ClientWebSocket.Close()
					}
				}
				wss.lock.Unlock()
			}
		}
	}()
}

func (wss *WebSocketServer) handleHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func (wss *WebSocketServer) handleEndPoint(w http.ResponseWriter, r *http.Request) {
	logger.Info("WebSocket Endpoint Hit")

	// Allow all origins
	wss.getUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// Upgrade connection to websocket
	ws, err := wss.getUpgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(err)
		return
	}

	// Setup firebase auth
	fb := firebaseAuth.Init()
	var idToken *auth.Token

	// Setup client
	client := &ClientObject{}
	defer func() {
		logger.Warning("Client disconnected: ", client.UID)
		wss.unregisterClient <- client
	}()

	// Setup controllersAuth
	controllers, controllersAuth := server.NewControllers()

	for {
		// Read message
		var msg Message
		if err := ws.ReadJSON(&msg); err != nil {
			logger.Error(err)
			break
		}
		data := msg.Data

		// Process messages without authentication
		if msgType, exists := controllers[msg.Type]; exists {
			if action, exists := msgType.(map[string]interface{})[msg.Action]; exists {
				if msg.Action != "ping" {
					logger.Info("Processing message without auth: ", msg.Type, " ", msg.Action)
				}
				action.(func(*websocket.Conn, map[string]interface{}) error)(ws, data)
				continue
			}
		}

		// Guard until authentication
		if idToken == nil {
			if msg.Type != "user" || msg.Action != "auth" {
				logger.Error("Must authenticate first")
				continue
			}

			idToken, err = firebaseAuth.ValidateToken(fb, data["token"].(string))
			if err != nil {
				logger.Error("Error verifying token: ", err)
				continue
			}

			ws.WriteJSON(map[string]string{"error": "null"})
			logger.Info("Authenticated: ", idToken.UID)

			// Register user in server
			client.UID = idToken.UID
			wss.registerClient <- client
		}

		// Process messages with authentication
		if msgType, exists := controllersAuth[msg.Type]; exists {
			if action, exists := msgType.(map[string]interface{})[msg.Action]; exists {
				logger.Info("Processing message with auth: ", idToken.UID, " ", msg.Type, " ", msg.Action)
				action.(func(*websocket.Conn, *auth.Token, map[string]interface{}) error)(ws, idToken, data)
				continue
			}
		}

		logger.Error("Invalid message type or action: ", msg.Type, " ", msg.Action)
	}
}
