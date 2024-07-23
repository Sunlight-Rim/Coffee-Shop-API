package delivery

import (
	"coffeeshop-api/internal/services/orders/model"
	"coffeeshop-api/pkg/errors"
	"coffeeshop-api/pkg/tools"

	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"
)

type handler struct {
	uc  model.IUsecase
	hub *Hub
}

func New(uc model.IUsecase, hub *Hub) *handler {
	return &handler{
		uc:  uc,
		hub: hub,
	}
}

func (h *handler) sseOrdersStatuses(c echo.Context) (err error) {
	// Parse request
	req := ordersStatusesReq(c)

	w := c.Response()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Register SSE client
	h.hub.registerClient <- req.UserID

	// Unregister SSE client after context was done
	go func() {
		<-c.Request().Context().Done()
		h.hub.unregisterClient <- req.UserID
	}()

	// Send events
	for status := range h.hub.clients[req.UserID] {
		logrus.WithField("user_id", req.UserID).Info("SSE sent order status update")

		if _, err := w.Write(append(status, []byte("\n")...)); err != nil {
			return errors.Wrap(err, "write to the client")
		}

		w.Flush()
	}

	return
}

func (h *handler) listOrders(c echo.Context) (err error) {
	var (
		req *model.ListOrdersReqDelivery
		res *model.ListOrdersResDelivery
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = listOrdersReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	orders, err := h.uc.ListOrders(&model.ListOrdersReqUsecase{
		UserID: req.UserID,
		Offset: req.Offset,
	})
	if err != nil {
		return errors.Wrap(err, "list orders")
	}

	res = &model.ListOrdersResDelivery{
		Orders: orders.Orders,
	}

	return
}

func (h *handler) getOrderInfo(c echo.Context) (err error) {
	var (
		req *model.GetOrderInfoReqDelivery
		res *model.GetOrderInfoResDelivery
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = getOrderInfoReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	order, err := h.uc.GetOrderInfo(&model.GetOrderInfoReqUsecase{
		UserID:  req.UserID,
		OrderID: req.OrderID,
	})
	if err != nil {
		return errors.Wrap(err, "get order info")
	}

	res = &model.GetOrderInfoResDelivery{
		Order: order.Order,
	}

	return
}

func (h *handler) createOrder(c echo.Context) (err error) {
	var (
		req *model.CreateOrderReqDelivery
		res *model.CreateOrderResDelivery
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = createOrderReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	order, err := h.uc.CreateOrder(&model.CreateOrderReqUsecase{
		UserID:  req.UserID,
		Address: req.Address,
		Items:   req.Items,
	})
	if err != nil {
		return errors.Wrap(err, "create order")
	}

	res = &model.CreateOrderResDelivery{
		OrderID: order.OrderID,
	}

	return
}

func (h *handler) cancelOrder(c echo.Context) (err error) {
	var (
		req *model.CancelOrderReqDelivery
		res *model.CancelOrderResDelivery
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = cancelOrderReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	order, err := h.uc.CancelOrder(&model.CancelOrderReqUsecase{
		UserID:  req.UserID,
		OrderID: req.OrderID,
	})
	if err != nil {
		return errors.Wrap(err, "cancel order")
	}

	res = &model.CancelOrderResDelivery{
		OrderID: order.OrderID,
	}

	return
}

func (h *handler) employeeCompleteOrder(c echo.Context) (err error) {
	var (
		req *model.EmployeeCompleteOrderReqDelivery
		res *model.EmployeeCompleteOrderResDelivery
	)

	// Send response
	defer func() { tools.SendResponse(c, res, err) }()

	// Parse request
	if req, err = employeeCompleteOrderReq(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	order, err := h.uc.EmployeeCompleteOrder(&model.EmployeeCompleteOrderReqUsecase{
		OrderID: req.OrderID,
	})
	if err != nil {
		return errors.Wrap(err, "complete order as employee")
	}

	res = &model.EmployeeCompleteOrderResDelivery{
		OrderID: order.OrderID,
	}

	// Send status to user SSE connection
	msg, err := easyjson.Marshal(model.OrdersStatusesResDelivery{
		OrderID:        order.OrderID,
		OrderCreatedAt: order.OrderCreatedAt,
		OrderStatus:    order.OrderStatus,
	})
	if err != nil {
		return errors.Wrap(err, "send order status to user SSE")
	}

	if _, ok := h.hub.clients[order.OrderCustomerID]; ok {
		h.hub.clients[order.OrderCustomerID] <- msg
	}

	return
}
