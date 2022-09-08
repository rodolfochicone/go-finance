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
	router.GET("/users/username/:username", server.GetUserByUsername)
	router.GET("/users/:id", server.GetUserByID)

	router.POST("/categories", server.CreateCategory)
	router.GET("/categories/:id", server.GetCategoryByID)
	router.GET("/categories", server.GetCategories)
	router.DELETE("/categories/:id", server.DeleteCategoryByID)
	router.PUT("/categories/:id", server.UpdateCategory)

	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts", server.GetAccounts)
	router.GET("/accounts/:id", server.GetAccountByID)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api has error": err.Error()}
}
