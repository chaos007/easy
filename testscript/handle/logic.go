package handle

import (
	"fmt"

	"github.com/chaos007/easy/data/pb"
	"github.com/chaos007/easy/testscript/player"

	"github.com/golang/protobuf/proto"
)

// DownPing DownPing
func DownPing(content proto.Message, p *player.Player) (ret proto.Message) {
	msg := content.(*pb.DownPing)
	fmt.Println("------DownPing:", msg)
	//p.Done("DownPing")
	return
}

// DownPlayerAction DownPlayerAction
func DownPlayerAction(content proto.Message, p *player.Player) (ret proto.Message) {
	// msg := content.(*pb.DownPlayerAction)
	// fmt.Println("------DownPlayerAction:", msg)
	// p.Done("DownPlayerAction")
	return
}

// GameFrame GameFrame
func GameFrame(content proto.Message, p *player.Player) (ret proto.Message) {
	// msg := content.(*pb.GameFrame)
	// fmt.Println("------GameFrame:", msg)
	// p.Done("GameFrame")
	return
}

// DownPlayerLogin DownPlayerLogin
func DownPlayerLogin(content proto.Message, p *player.Player) (ret proto.Message) {
	msg := content.(*pb.DownPlayerLogin)
	fmt.Println("------DownPlayerLogin:", msg)
	return
}

// DownJoinRoomToClient DownJoinRoomToClient
func DownJoinRoomToClient(content proto.Message, p *player.Player) (ret proto.Message) {
	msg := content.(*pb.DownJoinRoomToClient)
	fmt.Println("------DownJoinRoomToClient:", msg)
	return
}
