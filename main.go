package main

import (
	_ "homework/boot"
	_ "homework/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
