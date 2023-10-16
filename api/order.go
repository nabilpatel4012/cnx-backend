package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/nexpictora-pvt-ltd/cnx-backend/db/sqlc"
)

type createOrderRequest struct {
	OrderID     int64   `json:"order_id" binding:"required"`
	CustomerID  int     `json:"customer_id" binding:"required"`
	ServiceIDs  []int32 `json:"service_ids" binding:"required"`
	OrderStatus string  `json:"order_status" binding:"required"`
}

// func (server *Server) createOrder(ctx *gin.Context) {
// 	var req createOrderRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	// Fix: Add a check to ensure that the service IDs array is not empty.
// 	if len(req.ServiceIDs) == 0 {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "service_ids array cannot be empty"})
// 		return
// 	}
// 	serviceIDs := pq.Int64Array(req.ServiceIDs)
// 	var createdOrders []db.Order
// 	// Iterate through the service IDs and create individual orders
// 	for _, serviceID := range serviceIDs {

// 		createdOrder, err := server.store.CreateOrder(ctx, db.CreateOrderParams{
// 			OrderID:     int64(req.OrderID),
// 			UserID:      int32(req.CustomerID),
// 			ServiceIds:  int32(serviceID),
// 			OrderStatus: req.OrderStatus,
// 			// Set other fields as needed
// 		})
// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 			return
// 		}
// 		createdOrders = append(createdOrders, createdOrder)
// 	}
// 	ctx.JSON(http.StatusOK, createdOrders)
// }

// func (server *Server) createOrder(ctx *gin.Context) {
// 	var req createOrderRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	// Fix: Add a check to ensure that the service IDs array is not empty.
// 	if len(req.ServiceIDs) == 0 {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "service_ids array cannot be empty"})
// 		return
// 	}

// 	serviceIDs := pq.Int64Array(req.ServiceIDs)
// 	var createdOrdersWithServices []struct {
// 		ID                int64      `json:"id"`
// 		OrderID           int64      `json:"order_id"`
// 		UserID            int        `json:"user_id"`
// 		OrderStatus       string     `json:"order_status"`
// 		OrderStarted      time.Time  `json:"order_started"`
// 		OrderDelivered    bool       `json:"order_delivered"`
// 		OrderDeliveryTime time.Time  `json:"order_delivery_time"`
// 		Services          db.Service `json:"services"`
// 	}

// 	// Iterate through the service IDs and create individual orders
// 	for _, serviceID := range serviceIDs {
// 		createdOrder, err := server.store.CreateOrder(ctx, db.CreateOrderParams{
// 			OrderID:     int64(req.OrderID),
// 			UserID:      int32(req.CustomerID),
// 			ServiceIds:  int32(serviceID),
// 			OrderStatus: req.OrderStatus,
// 			// Set other fields as needed
// 		})
// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 			return
// 		}

// 		// Fetch service details
// 		service, err := server.store.GetService(ctx, int32(serviceID))
// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 			return
// 		}

// 		createdOrdersWithServices = append(createdOrdersWithServices, struct {
// 			ID                int64      `json:"id"`
// 			OrderID           int64      `json:"order_id"`
// 			UserID            int        `json:"user_id"`
// 			OrderStatus       string     `json:"order_status"`
// 			OrderStarted      time.Time  `json:"order_started"`
// 			OrderDelivered    bool       `json:"order_delivered"`
// 			OrderDeliveryTime time.Time  `json:"order_delivery_time"`
// 			Services          db.Service `json:"services"`
// 		}{
// 			ID:                int64(createdOrder.ID),
// 			OrderID:           createdOrder.OrderID,
// 			UserID:            int(createdOrder.UserID),
// 			OrderStatus:       createdOrder.OrderStatus,
// 			OrderStarted:      createdOrder.OrderStarted,
// 			OrderDelivered:    createdOrder.OrderDelivered,
// 			OrderDeliveryTime: createdOrder.OrderDeliveryTime,
// 			Services:          service,
// 		})
// 	}

// 	ctx.JSON(http.StatusOK, createdOrdersWithServices)
// }

type orderResponse struct {
	OrderID           int64     `json:"order_id"`
	UserID            int       `json:"user_id"`
	OrderStatus       string    `json:"order_status"`
	OrderStarted      time.Time `json:"order_started"`
	OrderDelivered    bool      `json:"order_delivered"`
	OrderDeliveryTime time.Time `json:"order_delivery_time"`
	Services          []struct {
		ServiceID    int64  `json:"service_id"`
		ServiceName  string `json:"service_name"`
		ServicePrice int    `json:"service_price"`
		ServiceImage string `json:"service_image"`
	} `json:"services"`
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

	var services []struct {
		ServiceID    int64  `json:"service_id"`
		ServiceName  string `json:"service_name"`
		ServicePrice int    `json:"service_price"`
		ServiceImage string `json:"service_image"`
	}
	var createdOrders []db.Order
	// Fetch service details for the response
	for _, serviceID := range req.ServiceIDs {

		createdOrder, err := server.store.CreateOrder(ctx, db.CreateOrderParams{
			OrderID:     req.OrderID,
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
		// Fetch service details for the response
		service, err := server.store.GetService(ctx, int32(serviceID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		services = append(services, struct {
			ServiceID    int64  `json:"service_id"`
			ServiceName  string `json:"service_name"`
			ServicePrice int    `json:"service_price"`
			ServiceImage string `json:"service_image"`
		}{
			ServiceID:    int64(service.ServiceID),
			ServiceName:  service.ServiceName,
			ServicePrice: int(service.ServicePrice),
			ServiceImage: service.ServiceImage,
		})
	}

	// Construct the response object
	response := struct {
		OrderID           int64     `json:"order_id"`
		UserID            int       `json:"user_id"`
		OrderStatus       string    `json:"order_status"`
		OrderStarted      time.Time `json:"order_started"`
		OrderDelivered    bool      `json:"order_delivered"`
		OrderDeliveryTime time.Time `json:"order_delivery_time"`
		Services          []struct {
			ServiceID    int64  `json:"service_id"`
			ServiceName  string `json:"service_name"`
			ServicePrice int    `json:"service_price"`
			ServiceImage string `json:"service_image"`
		} `json:"services"`
	}{
		OrderID:           req.OrderID,
		UserID:            req.CustomerID,
		OrderStatus:       req.OrderStatus,
		OrderStarted:      createdOrders[0].OrderStarted, // Assuming you want the order_started time of the first created order
		OrderDelivered:    createdOrders[0].OrderDelivered,
		OrderDeliveryTime: createdOrders[0].OrderDeliveryTime,
		Services:          services,
	}

	ctx.JSON(http.StatusOK, response)
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

type listOrdersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listOrders(ctx *gin.Context) {
	var req listOrdersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListOrdersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	services, err := server.store.ListOrders(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, services)
}
