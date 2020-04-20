package controller

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	objectid := "abc"
	this.Data["json"] = "{\"ObjectId\":\"" + objectid + "\"}"
	this.ServeJSON()
}
