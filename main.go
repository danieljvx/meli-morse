package main

import (
	"github.com/danieljvx/meli-morse/app"
	_ "github.com/danieljvx/meli-morse/docs"
	"os"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email villanueva.danielx@gmail.com
// @BasePath /
func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}
	app := app.App()
	app.Listen(":" + port)
}
