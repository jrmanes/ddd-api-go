package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	mooc "github.com/jrmanes/ddd-api-go/internal"
	"github.com/jrmanes/ddd-api-go/internal/platform/server/handler/courses"
	"github.com/jrmanes/ddd-api-go/internal/platform/server/handler/health"
)

type Server struct {
	httpAddr string
	engine *gin.Engine

	// dependencies
	courseRepository mooc.CourseRepository
}

// New initialization server Gin
func New(host string, port uint, courseRepository mooc.CourseRepository) Server {
	srv := Server{
		httpAddr: fmt.Sprintf(host,port),
		engine: gin.New(),

		courseRepository: courseRepository,
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
	s.engine.POST("/courses", courses.CreateHandler(s.courseRepository))
}