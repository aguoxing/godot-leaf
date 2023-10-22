package gate

import (
	"server/src/server/game"
	"server/src/server/msg"
)

func init() {
	// 这里指定消息 Hello 路由到 game 模块
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&msg.PlayerRequest{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.PlayerResponse{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.NewGame{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.StartGame{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.ExitGame{}, game.ChanRPC)
}
