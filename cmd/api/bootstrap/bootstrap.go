package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	mooc "github.com/jrmanes/ddd-api-go/internal"
	"github.com/jrmanes/ddd-api-go/internal/creating"
	"github.com/jrmanes/ddd-api-go/internal/increasing"
	"github.com/jrmanes/ddd-api-go/internal/platform/bus/inmemory"
	"github.com/jrmanes/ddd-api-go/internal/platform/server"
	"github.com/jrmanes/ddd-api-go/internal/platform/storage/mysql"
	"github.com/kelseyhightower/envconfig"
)

// Run initialize the server, if st goes wrong, return an error to main
func Run() error {
	var cfg config
	err := envconfig.Process("MOOC", &cfg)
	if err != nil {
		return err
	}

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)

	courseRepository := mysql.NewCourseRepository(db, cfg.DbTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository, eventBus)
	increasingCourseCounterService := increasing.NewCourseCounterService()

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	eventBus.Subscribe(
		mooc.CourseCreatedEventType,
		creating.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseCounterService),
	)

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus)
	return srv.Run(ctx)
}

// config sets the variables in order to get them from env and default values
type config struct {
	// Server configuration
	Host            string        `default:"0.0.0.0"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser    string        `default:"api-ddd"`
	DbPass    string        `default:"api-ddd"`
	DbHost    string        `default:"localhost"`
	DbPort    uint          `default:"3306"`
	DbName    string        `default:"api-ddd"`
	DbTimeout time.Duration `default:"5s"`
}
