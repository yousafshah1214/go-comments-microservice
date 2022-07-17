package main

import (
	"fmt"

	"github.com/yousafshah1214/go-comments-microservice/internal/comment"
	"github.com/yousafshah1214/go-comments-microservice/internal/database"
	transportHTTP "github.com/yousafshah1214/go-comments-microservice/internal/transport/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up our App")

	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.Migrate(db)
	if err != nil {
		return err
	}
	service := comment.NewService(db)
	handler := transportHTTP.NewHandler(service)
	err = handler.SetupRoutes()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting REST API")
		fmt.Println(err)
	}
}
