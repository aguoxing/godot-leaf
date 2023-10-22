package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"server/src/server/msg"
)

var rooms = make(map[string][]*msg.PlayerInfo)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.PlayerRequest{}, handlePlayerReq)
	handler(&msg.NewGame{}, handleNewGame)
	handler(&msg.StartGame{}, handleStartGame)
	handler(&msg.ExitGame{}, handleExitGame)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlePlayerReq(args []interface{}) {
	m := args[0].(*msg.PlayerRequest)
	// 消息的发送者
	u := args[1].(gate.Agent)
	u.SetUserData(m.PlayerId)

	// 输出收到的消息的内容
	log.Debug("%v move %v", m.PlayerId, m.Vel)

	// 给发送者回应一个 Hello 消息
	//u.WriteMsg(&msg.PlayerResponse{
	//	Data: m.Name,
	//})

	// 广播
	for a := range agents {
		d, _ := a.UserData().(string)
		if d != m.PlayerId {
			a.WriteMsg(&msg.PlayerResponse{
				PlayerId: m.PlayerId,
				Vel:      m.Vel,
			})
		}
	}
}

func handleNewGame(args []interface{}) {
	m := args[0].(*msg.NewGame)
	// 消息的发送者
	u := args[1].(gate.Agent)
	_ = u

	// 输出收到的消息的内容
	//log.Debug("roomId: %v", m.RoomId)

	var res = ""
	// 加入or创建room
	if playerList, ok := rooms[m.RoomId]; ok {
		playerInfo := &msg.PlayerInfo{
			PlayerId: m.PlayerId,
			X:        50,
			Y:        50,
		}
		rooms[m.RoomId] = append(rooms[m.RoomId], playerInfo)
		res = m.PlayerId + ": join[" + m.RoomId + "]room"
		log.Debug(res)
	} else {
		playerList = []*msg.PlayerInfo{
			{
				PlayerId: m.PlayerId,
				X:        100,
				Y:        100,
			},
		}
		rooms[m.RoomId] = playerList
		res = m.PlayerId + ": create[" + m.RoomId + "]room"
		log.Debug(res)
	}

	// 给发送者回应一个 success 消息
	u.WriteMsg(&msg.PlayerResponse{
		Msg: res,
	})
}

func handleStartGame(args []interface{}) {
	m := args[0].(*msg.StartGame)
	// 消息的发送者
	u := args[1].(gate.Agent)
	_ = u

	// 输出收到的消息的内容
	log.Debug("roomId[%v] start game", m.RoomId)

	// 广播
	for a := range agents {
		a.WriteMsg(&msg.StartGame{
			RoomId:     m.RoomId,
			PlayerList: rooms[m.RoomId],
		})
	}

	fmt.Print(rooms[m.RoomId])
}

func handleExitGame(args []interface{}) {
	m := args[0].(*msg.ExitGame)
	// 消息的发送者
	u := args[1].(gate.Agent)
	_ = u

	rooms[m.RoomId] = removeElementByVal(rooms[m.RoomId], m.PlayerId)
	playerCount := len(rooms[m.RoomId])

	if playerCount == 0 {
		log.Debug("destroy room[%v]", m.RoomId)
		delete(rooms, m.RoomId)
	}

	log.Debug("player[%v] exit room[%v] playerCount[%v]", m.PlayerId, m.RoomId, playerCount)

	// 广播
	for a := range agents {
		a.WriteMsg(&msg.ExitGame{
			RoomId:   m.RoomId,
			PlayerId: m.PlayerId,
		})
	}
}

// 删除指定下标的元素
func removeElementByIdx(arr []string, index int) []string {
	// 判断下标是否越界
	if index < 0 || index >= len(arr) {
		return arr
	}

	// 使用切片删除元素
	return append(arr[:index], arr[index+1:]...)
}

func removeElementByVal(arr []*msg.PlayerInfo, playerId string) []*msg.PlayerInfo {
	for i := 0; i < len(arr); i++ {
		if arr[i].PlayerId == playerId {
			arr = append(arr[:i], arr[i+1:]...)
			break
		}
	}

	return arr
}

func removeElementByVal1(arr []string, val string) []string {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			arr = append(arr[:i], arr[i+1:]...)
			break
		}
	}

	return arr
}
