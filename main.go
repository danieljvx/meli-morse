package main

import (
	"github.com/danieljvx/meli-morse/app"
)

func main() {
	app := app.App()
	app.Listen(":8000")
}
