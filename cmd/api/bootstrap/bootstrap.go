package bootstrap

import "github.com/jrmanes/ddd-api-go/internal/platform/server"

const (
	host = "localhost"
	port = 8080
)

// Run initialize the server, if st goes wrong, return an error to main
func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}