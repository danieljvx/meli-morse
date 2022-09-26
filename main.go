package main

import (
	"github.com/danieljvx/meli-morse/app"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8000"
	}
	app := app.App()
	app.Listen(port)
}
