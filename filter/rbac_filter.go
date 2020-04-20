package filter

import (
	"fmt"

	"github.com/astaxie/beego/context"
)

var RBACFilter = func(ctx *context.Context) {
	fmt.Println("=== RBACFilter")
}
