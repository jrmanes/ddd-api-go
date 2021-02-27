package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jrmanes/ddd-api-go/internal/platform/server/handler/health"
)

type Server struct {
	httpAddr string
	engine *gin.Engine
}

// New initialization server Gin
func New(host string, port uint) Server {
	srv := Server{
		httpAddr: fmt.Sprintf(host,port),
		engine: gin.New(),
	}

	srv.registerRoutes()

	return srv
}

func (s *Server) Run() error  {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

// registerRoutes is the place that we defined all the routes
func (s *Server) registerRoutes()  {
	s.engine.GET("/health", health.CheckHandler())
}