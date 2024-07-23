package delivery

import (
	"github.com/sirupsen/logrus"
)

// Server-Sent Events management hub.
type Hub struct {
	// Add connection
	registerClient chan uint64
	// Close connection
	unregisterClient chan uint64
	// Connections registry
	clients map[uint64](chan []byte)
}

func NewHubSSE() *Hub {
	return &Hub{
		registerClient:   make(chan uint64),
		unregisterClient: make(chan uint64),
		clients:          make(map[uint64](chan []byte)),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case userID := <-h.registerClient:
			logrus.WithField("user_id", userID).Info("SSE client connected")
			h.clients[userID] = make(chan []byte, 20)

		case userID := <-h.unregisterClient:
			logrus.WithField("user_id", userID).Info("SSE client disconnected")
			close(h.clients[userID])
			delete(h.clients, userID)
		}
	}
}
