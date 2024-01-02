package api

import (
	db "github.com/ayo6706/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// server serves http requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// New  server creates  a new http server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
