package wss

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	UID  string
	Conn *websocket.Conn
}

type WSS struct {
	// Clients          map[*Client]bool
	Clients sync.Map
	// ClientUIDMap     map[string]*Client
	ClientUIDMap     sync.Map
	RegisterClient   chan *Client
	UnregisterClient chan *Client

	GetUpgrader websocket.Upgrader

	Lock sync.Mutex
}
