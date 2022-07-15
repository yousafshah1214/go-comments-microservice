package main

import (
	"fmt"

	transportHTTP "github.com/yousafshah1214/go-comments-microservice/internal/transport/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up our App")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()
	return nil
}

func main() {

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting REST API")
		fmt.Println(err)
	}
}
