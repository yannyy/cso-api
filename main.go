package main

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	// Create service
	service := web.NewService(
		web.Name("api-gateway"),
	)

	service.Init()

	// Register Handler
	service.Handle("/", beego.BeeApp.Handlers)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
