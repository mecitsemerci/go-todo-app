package main

import (
	"fmt"
	"github.com/mecitsemerci/clean-go-todo-api/app"
	"os"
)

// @title Todo API
// @version 1.0.0
// @description This is a sample todo restful api server.
// @host localhost:8080
// @BasePath /
func main() {

	// Run app
	if err := app.Init().Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
