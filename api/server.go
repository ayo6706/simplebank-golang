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

	server.router = router
	return server
}
