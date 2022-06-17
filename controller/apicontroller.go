package controller

import "github.com/xueqianLu/nodeproxy/ethrpc"

type RPCController struct {
	BaseController
}

func (e *RPCController) PeerInfo() {
	var result, err = ethrpc.NetPeers()
	if err != nil {
		e.ResponseInfo(500, err, result)
		return
	}
	e.ResponseInfo(200, nil, result)
}

