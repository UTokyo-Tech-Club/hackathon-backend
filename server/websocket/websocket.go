package websocket

import (
	"fmt"
	firebaseAuth "hackathon-backend/firebase"
	"hackathon-backend/server"
	"sync"

	"hackathon-backend/utils/logger"
	"net/http"

	wss "hackathon-backend/server/websocketServer"

	"firebase.google.com/go/auth"
	"github.com/gorilla/websocket"
)

var ws *wss.WSS

func Init() {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws = &wss.WSS{
		// Clients:          make(map[*wss.Client]bool),
		Clients: sync.Map{},
		// ClientUIDMap:     make(map[string]*wss.Client),
		ClientUIDMap:     sync.Map{},
		RegisterClient:   make(chan *wss.Client),
		UnregisterClient: make(chan *wss.Client),

		GetUpgrader: upgrader,

		Lock: sync.Mutex{},
	}

	setupRouter()
	setupEventListeners()
}

func setupRouter() {
	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/ws", handleEndPoint)
}

func setupEventListeners() {

	// Client Registration
	go func() {
		for {
			select {
			case client := <-ws.RegisterClient:
				// ws.Lock.Lock()
				// ws.Clients[client] = true
				// ws.ClientUIDMap[client.UID] = client
				// ws.Lock.Unlock()
				ws.Clients.Store(client, true)
				ws.ClientUIDMap.Store(client.UID, client)
			case client := <-ws.UnregisterClient:
				// ws.Lock.Lock()
				// if _, ok := ws.Clients[client]; ok {
				// 	delete(ws.Clients, client)
				// 	delete(ws.ClientUIDMap, client.UID)
				// 	if client.Conn != nil {
				// 		client.Conn.Close()
				// 	}
				// }
				// ws.Lock.Unlock()
				ws.Clients.Delete(client)
				ws.ClientUIDMap.Delete(client.UID)
			}
		}
	}()
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

// First establish HTTP connection, then upgrade to WebSocket
// Communication through WebSocket tunnel is handled with *websocket.Conn for funtionalities allowing anonymous access
// Functionalities requiring authentication are handled with *wss.WSS that enables broadcasting
func handleEndPoint(w http.ResponseWriter, r *http.Request) {
	logger.Info("WebSocket Endpoint Hit")

	// Allow all origins
	ws.GetUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// Upgrade connection to websocket
	socket, err := ws.GetUpgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(err)
		return
	}

	// Setup firebase auth
	fb := firebaseAuth.Init()
	var idToken *auth.Token

	// Setup controllers
	controllers, controllersAuth := server.NewControllers()

	for {
		// Read message
		var msg Message
		if err := socket.ReadJSON(&msg); err != nil {
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
				action.(func(*websocket.Conn, map[string]interface{}) error)(socket, data)
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

			socket.WriteJSON(map[string]string{"error": "null"})
			logger.Info("Authenticated: ", idToken.UID)

			// Add client to WebSocket server
			client := &wss.Client{}
			defer func() {
				logger.Warning("Client disconnected: ", client.UID)
				ws.UnregisterClient <- client
			}()
			client.UID = idToken.UID
			client.Conn = socket

			ws.RegisterClient <- client

			// Wait until client is registered to WebSocket server
			for {
				// if _, ok := ws.ClientUIDMap[idToken.UID]; ok {
				// 	break
				// }
				if _, ok := ws.ClientUIDMap.Load(idToken.UID); ok {
					break
				}
			}

		}

		// Process messages with authentication
		if msgType, exists := controllersAuth[msg.Type]; exists {
			if action, exists := msgType.(map[string]interface{})[msg.Action]; exists {
				logger.Info("Processing message with auth: ", idToken.UID, " ", msg.Type, " ", msg.Action)
				action.(func(*wss.WSS, *auth.Token, map[string]interface{}) error)(ws, idToken, data)
				continue
			}
		}

		logger.Error("Invalid message type or action: ", msg.Type, " ", msg.Action)
		socket.WriteJSON(map[string]interface{}{"error": "invalid type or action"})
	}
}
