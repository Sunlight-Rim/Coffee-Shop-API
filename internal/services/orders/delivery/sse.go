package delivery

import (
	"coffeeshop-api/internal/services/orders/model"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tools"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var (
	// Storing SSE clients by user ID.
	sseClients = make(map[uint64](chan []byte))
	mu         sync.Mutex
)

func (h *handler) ordersStatuses(c echo.Context) (err error) {
	var req *model.OrdersStatusesReqDelivery

	// Parse request
	if req, err = ordersStatusesReq(c); err != nil {
		err = errors.Wrap(err, "request")
		tools.SendResponse(c, nil, err)
		return
	}

	// Register SSE client
	w := c.Response()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	mu.Lock()
	sseClients[req.UserID] = make(chan []byte, 50)
	mu.Unlock()

	logrus.WithField("user_id", req.UserID).Info("New SSE client connected")

	// Listen SSE events
	for {
		select {
		case <-c.Request().Context().Done():
			logrus.WithField("user_id", req.UserID).Info("SSE client disconnected")
			return

		case status := <-sseClients[req.UserID]:
			logrus.WithField("user_id", req.UserID).Info("SSE sent order status update")

			status = append(status, []byte("\n")...)
			if _, err := w.Write(status); err != nil {
				return errors.Wrap(err, "write to the client")
			}

			w.Flush()
		}
	}
}
