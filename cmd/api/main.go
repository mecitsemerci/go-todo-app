package main

import (
	"log"
	"os"

	_ "github.com/mecitsemerci/go-todo-app/docs"
	"github.com/mecitsemerci/go-todo-app/internal/api"
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
		log.Printf("%s", err.Error())
		os.Exit(1)
	}
}
