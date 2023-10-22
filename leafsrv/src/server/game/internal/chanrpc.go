package internal

import (
	"github.com/name5566/leaf/gate"
)

// 存储所有连接 可用于广播
var agents = make(map[gate.Agent]struct{})

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	agents[a] = struct{}{}
	//_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	delete(agents, a)
	//_ = a
}
