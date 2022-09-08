package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/rodolfochicone/go-finance/db/sqlc"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.CreateUser)
	router.GET("/users/username/:username", server.GetUser)
	router.GET("/users/:id", server.GetUserByID)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api has error": err.Error()}
}
