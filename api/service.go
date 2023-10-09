package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/nexpictora-pvt-ltd/cnx-backend/db/sqlc"
)

type createServiceRequest struct {
	ServiceName  string `json:"service_name" binding:"required"`
	ServicePrice int64  `json:"service_price" binding:"required"`
}

func (server *Server) createService(ctx *gin.Context) {
	var req createServiceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateServiceParams{
		ServiceName:  req.ServiceName,
		ServicePrice: req.ServicePrice,
	}

	service, err := server.store.CreateService(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, service)
}

type getServiceRequest struct {
	ServiceID int64 `uri:"service_id" binding:"required,min=1"`
}

func (server *Server) getService(ctx *gin.Context) {
	var req getServiceRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	service, err := server.store.GetService(ctx, int32(req.ServiceID))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, service)
}

type listServicesRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listServices(ctx *gin.Context) {
	var req listServicesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListLimitedServicesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	services, err := server.store.ListLimitedServices(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, services)
}

type updateServiceRequest struct {
	ServiceID    int64  `uri:"service_id" binding:"required,min=1"`
	ServiceName  string `json:"service_name" binding:"required"`
	ServicePrice int64  `json:"service_price" binding:"required"`
}

func (server *Server) updateService(ctx *gin.Context) {
	var req updateServiceRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	updateParams := db.UpdateServiceParams{
		ServiceID:    int32(req.ServiceID),
		ServiceName:  req.ServiceName,
		ServicePrice: req.ServicePrice,
	}

	service, err := server.store.UpdateService(ctx, updateParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, service)
}
