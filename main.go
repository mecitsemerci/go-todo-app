package main

import (
	"github.com/mecitsemerci/clean-go-todo-api/app"
	"log"
)

// @title Todo API
// @version 1.0.0
// @description This is a sample todo restful api server.
// @host localhost:8080
// @BasePath /
func main() {

	// Run app
	if err := app.Init().Run(); err != nil {
		log.Fatal(err)
		return
	}
}
