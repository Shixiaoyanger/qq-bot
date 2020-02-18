package view

import (
	"net/http"
)

func ListenMessage() {
	// Or if you love WebSocket Reverse
	// updates := bot.ListenForWebSocket(u)
	go http.ListenAndServe("0.0.0.0:8443", nil)
}
