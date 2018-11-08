package main

import (
	"web-restful-golang/app"
	"web-restful-golang/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
