package main

import (
	"github.com/astaxie/beego"
	_ "github.com/xueqianLu/nodeproxy/ethrpc"
	_ "github.com/xueqianLu/nodeproxy/routers"
)

func main() {
	beego.Run()
}

