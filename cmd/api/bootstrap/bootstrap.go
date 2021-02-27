package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/jrmanes/ddd-api-go/internal/platform/server"
	"github.com/jrmanes/ddd-api-go/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "ddd-api-go"
	dbPass = "ddd-api-go"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "ddd-api-go"
)

// Run initialize the server, if st goes wrong, return an error to main
func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser,dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}