package main

import (
	auth "cso/auth/proto"
	"log"

	"github.com/astaxie/beego"
	bctx "github.com/astaxie/beego/context"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"
)

type Server struct {
	AuthBackend auth.AuthService
}

func (s *Server) Message(ctx *bctx.Context) {
	ctx.Output.JSON("abc", false, true)
}

func main() {
	// Create service
	service := web.NewService(
		web.Name("cso.api-gateway"),
	)

	service.Init()

	// Create RESTful handler
	server := new(Server)
	server.AuthBackend = auth.NewAuthService("cso.auth", client.DefaultClient)

	beego.Get("/message", server.Message)

	// Register Handler
	service.Handle("/", beego.BeeApp.Handlers)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
