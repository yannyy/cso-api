package filter

import (
	"fmt"

	"github.com/astaxie/beego/context"
)

var AuthFilter = func(ctx *context.Context) {
	fmt.Println("=== AuthFilter")
}
