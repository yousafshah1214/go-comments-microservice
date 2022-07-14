package main

import "fmt"

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up our App")
	return nil
}

func main() {

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting REST API")
		fmt.Println(err)
	}
}
