package api

import (
	db "github.com/devvyky/logistics/db/sqlc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/products", server.listPackSizes)
	router.GET("/products/:id", server.getPackSize)
	router.POST("/products", server.createPackSize)
	router.PUT("/products/:id", server.updatePackSize)
	router.DELETE("/products/:id", server.deletePackSize)

	router.POST("/orders", server.createOrder)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
