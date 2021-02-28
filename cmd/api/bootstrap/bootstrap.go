package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jrmanes/ddd-api-go/internal/creating"
	"github.com/jrmanes/ddd-api-go/internal/platform/bus/inmemory"
	"github.com/jrmanes/ddd-api-go/internal/platform/server"
	"github.com/jrmanes/ddd-api-go/internal/platform/storage/mysql"
)

const (
	host            = "localhost"
	port            = 8080
	shutdownTimeout = 10 * time.Second

	dbUser    = "api-go"
	dbPass    = "api-go"
	dbHost    = "localhost"
	dbPort    = "3306"
	dbName    = "api-go"
	dbTimeout = 5 * time.Second
)

// Run initialize the server, if st goes wrong, return an error to main
func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)

	courseRepository := mysql.NewCourseRepository(db, dbTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository, eventBus)

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}
