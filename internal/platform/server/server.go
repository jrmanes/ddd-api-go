package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jrmanes/ddd-api-go/internal/platform/server/handler/courses"
	"github.com/jrmanes/ddd-api-go/internal/platform/server/handler/health"
	"github.com/jrmanes/ddd-api-go/kit/command"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// dependencies
	commandBus command.Bus
}

// New initialization server Gin
func New(host string, port uint, commandBus command.Bus) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		commandBus: commandBus,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.commandBus))
}
