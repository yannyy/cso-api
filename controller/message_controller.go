package controller

import (
	"context"
	"cso/api-gateway/backend"
	"fmt"

	"github.com/astaxie/beego"
)

type MessageController struct {
	beego.Controller
}

func (this *MessageController) Message() {
	resp, err := backend.Message(context.Background())
	if err != nil {
		fmt.Printf("%+v", err)
	}

	this.Data["json"] = "{\"Message\":\"" + resp.GetMessage() + "\"}"
	this.ServeJSON()
}
