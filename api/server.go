package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/nexpictora-pvt-ltd/cnx-backend/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.GET("/users/:user_id", server.getUser)
	router.GET("/users", server.listUser)

	router.POST("/services", server.createService)
	router.GET("/services/:service_id", server.getService)
	router.GET("/services", server.listServices)
	router.PUT("/services/:service_id", server.updateService)

	server.router = router
	return server
}

// Start the http server on the input/specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
