package main

import (
	"context"
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
	resp, err := s.AuthBackend.Message(context.TODO(), &auth.Empty{})
	if err != nil {
		ctx.Output.SetStatus(500)
		ctx.Output.JSON(err, false, true)
	}
	log.Println(resp)
	ctx.Output.JSON(resp.Message, false, true)
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
