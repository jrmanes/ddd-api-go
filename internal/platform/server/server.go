package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jrmanes/ddd-api-go/internal/creating"
	"github.com/jrmanes/ddd-api-go/internal/platform/server/handler/courses"
	"github.com/jrmanes/ddd-api-go/internal/platform/server/handler/health"
)

type Server struct {
	httpAddr string
	engine *gin.Engine

	// dependencies
	creatingCourseService creating.CourseService
}

// New initialization server Gin
func New(host string, port uint, creatingCourseService creating.CourseService) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		creatingCourseService: creatingCourseService,
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
	s.engine.POST("/courses", courses.CreateHandler(s.creatingCourseService))
}