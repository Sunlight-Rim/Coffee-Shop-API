package delivery

import (
	"sync"

	logger "github.com/sirupsen/logrus"
)

type Hub struct {
	// Connections registry
	clients map[uint64](chan []byte)
	mu      sync.Mutex
}

func NewHubSSE() *Hub {
	return &Hub{
		clients: make(map[uint64](chan []byte)),
	}
}

func (h *Hub) registerClient(userID uint64) {
	logger.WithField("user_id", userID).Info("SSE client connected")

	h.mu.Lock()
	defer h.mu.Unlock()

	h.clients[userID] = make(chan []byte, 20)
}

func (h *Hub) unregisterClient(userID uint64) {
	logger.WithField("user_id", userID).Info("SSE client disconnected")

	h.mu.Lock()
	defer h.mu.Unlock()

	close(h.clients[userID])
	delete(h.clients, userID)
}
