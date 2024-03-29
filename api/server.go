package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/nexpictora-pvt-ltd/cnx-backend/db/sqlc"
	"github.com/nexpictora-pvt-ltd/cnx-backend/token"
	"github.com/nexpictora-pvt-ltd/cnx-backend/util"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	//By uncommenting this line and commenting above line we can use JWT token as access token as both of them use same maker interfaces
	// tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// Swagger Documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Anyone can use or create user or login routes
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	//Here the endpoint is used to renew access token for user session
	router.POST("/tokens/renew_access", server.renewAccessToken)
	// Here we are grouping all the routes and making them protected
	authRoutes := router.Group("/").Use(authMiddleWare(server.tokenMaker))
	// These Routes should be protected as not everyone should have access to it
	authRoutes.GET("/users/:user_id", server.getUser)
	router.GET("/users", server.listUser)

	// Admin can Create/Add service
	router.POST("/services", server.createService)

	// User Endpoints
	authRoutes.GET("/services/:service_id", server.getService)
	authRoutes.GET("/services/all", server.listServices)

	//Admin Endpoints
	router.GET("/services/preview", server.listServices)
	router.GET("/services", server.listLimitedServices)
	router.PUT("/services/:service_id", server.updateService)
	router.DELETE("/services/:service_id", server.deleteService)

	authRoutes.POST("/orders", server.createOrder)
	authRoutes.PUT("/orders", server.updateOrderStatus)
	authRoutes.PUT("/orders/update", server.updateOrderDelivered)
	authRoutes.GET("/orders/:order_id", server.getOrder)
	authRoutes.GET("/orders", server.listOrders)

	server.router = router
}

// Start the http server on the input/specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
