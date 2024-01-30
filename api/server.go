package api

import (
	"fmt"

	db "github.com/ayo6706/simplebank/db/sqlc"
	"github.com/ayo6706/simplebank/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// server serves http requests
type Server struct {
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// New  server creates  a new http server and setup routing
func NewServer(store db.Store) *Server {
	tokenMaker, err := token.NewPasetoMaker("")
	if err != nil {
		return nil, fmt.Errorf("cannot crea")
	}
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	router.POST("/users", server.createUser)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
