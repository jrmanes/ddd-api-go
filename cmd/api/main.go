package main

import (
	"log"

	"github.com/jrmanes/ddd-api-go/cmd/api/bootstrap"
)

func main()  {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
