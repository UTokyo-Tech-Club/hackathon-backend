package websocket

type Message struct {
	Type   string                 `json:"type"`
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data,omitempty"`
}
