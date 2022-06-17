package routers

import (
	"github.com/astaxie/beego"
	"github.com/xueqianLu/nodeproxy/controller"
)

func init() {
	ns := beego.NewNamespace("/nodeproxy",
		beego.NSNamespace("api",
			//用户信息
			beego.NSRouter("/peerinfo", &controller.RPCController{}, "get:PeerInfo"),
		),
	)
	beego.AddNamespace(ns)
}
