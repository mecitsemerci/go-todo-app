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
	engine := api.Setup()

	if err := engine.Run(); err != nil {
		log.Printf("%s", err.Error())
		os.Exit(1)
	}
	defer engine.Close()
}
