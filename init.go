package main

import (
	"cso/api-gateway/controller"
	"cso/api-gateway/filter"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/message", &controller.MessageController{}, "*:Message")
	beego.Router("/resources/company", &controller.CompanyController{})
	beego.Router("/", &controller.IndexController{})
}

func init() {
	beego.InsertFilter("/resources/*", beego.BeforeRouter, filter.AuthFilter)
	beego.InsertFilter("/resources/*", beego.BeforeRouter, filter.RBACFilter)
}
