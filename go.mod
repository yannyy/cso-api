module cso/api-gateway

go 1.14

require (
	cso/auth v0.0.0
	github.com/astaxie/beego v1.12.1
	github.com/micro/go-micro v1.18.0
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
)

replace cso/auth => ../auth
