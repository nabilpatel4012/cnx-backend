package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/nexpictora-pvt-ltd/cnx-backend/db/sqlc"
)

type createOrderRequest struct {
	OrderID     int64         `json:"order_id" binding:"required"`
	CustomerID  int           `json:"customer_id" binding:"required"`
	ServiceIDs  pq.Int64Array `json:"service_ids" binding:"required"`
	OrderStatus string        `json:"order_status" binding:"required"`
}

func (server *Server) createOrder(ctx *gin.Context) {
	var req createOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Fix: Add a check to ensure that the service IDs array is not empty.
	if len(req.ServiceIDs) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "service_ids array cannot be empty"})
		return
	}
	serviceIDs := pq.Int64Array(req.ServiceIDs)
	var createdOrders []db.Order
	// Iterate through the service IDs and create individual orders
	for _, serviceID := range serviceIDs {

		createdOrder, err := server.store.CreateOrder(ctx, db.CreateOrderParams{
			OrderID:     int64(req.OrderID),
			UserID:      int32(req.CustomerID),
			ServiceIds:  int32(serviceID),
			OrderStatus: req.OrderStatus,
			// Set other fields as needed
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		createdOrders = append(createdOrders, createdOrder)
	}
	ctx.JSON(http.StatusOK, createdOrders)
}

type updateOrderStatusRequest struct {
	OrderID     int64  `json:"order_id" binding:"required"`
	OrderStatus string `json:"order_status" binding:"required"`
}

func (server *Server) updateOrderStatus(ctx *gin.Context) {
	var req updateOrderStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	updateOrderStatusParam := db.UpdateOrderStatusParams{
		OrderID:     req.OrderID,
		OrderStatus: req.OrderStatus,
	}

	orderStatus, err := server.store.UpdateOrderStatus(ctx, updateOrderStatusParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, orderStatus)
}

type updateOrderDeliveredRequest struct {
	OrderID           int64     `json:"order_id" binding:"required"`
	OrderDelivered    bool      `json:"order_delivered" binding:"required"`
	OrderDeliveryTime time.Time `json:"order_delivery_time" binding:"required"`
}

func (server *Server) updateOrderDelivered(ctx *gin.Context) {
	var req updateOrderDeliveredRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	updateOrderDeliveryParam := db.UpdateOrderDeliveryParams{
		OrderID:           req.OrderID,
		OrderDelivered:    req.OrderDelivered,
		OrderDeliveryTime: req.OrderDeliveryTime,
	}

	orderStatus, err := server.store.UpdateOrderDelivery(ctx, updateOrderDeliveryParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, orderStatus)
}

type getOrderRequest struct {
	OrderID int64 `uri:"order_id" binding:"required"`
}

func (server *Server) getOrder(ctx *gin.Context) {
	var req getOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	order, err := server.store.GetOrder(ctx, int64(req.OrderID))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, order)

}
