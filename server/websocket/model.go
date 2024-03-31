package websocket

type Message struct {
	Type   string `json:"type"`
	Action string `json:"action"`
	Data   string `json:"data,omitempty"`
}
