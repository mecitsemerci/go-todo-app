package main

import (
	"log"

	_ "github.com/mecitsemerci/go-todo-app/docs"
	"github.com/mecitsemerci/go-todo-app/internal/app/api"
)

// @title Todo API
// @version 1.0.0
// @description This is a sample todo restful api server.
// @host localhost:8080
// @BasePath /
func main() {
	app, err := api.NewAppServer()

	if err != nil {
		log.Fatalf("AppServer Error: %s", err.Error())
	}
	defer app.Close()

	if err := app.Start(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
