package handle

import (
	"github.com/chaos007/easy/game/handle/controller"
	"github.com/chaos007/easycome/session/grpc"
)

// Init Init
func Init() {
	grpc.RegisterProtocol("pb.TestToGame", controller.TestToGame)
	grpc.RegisterProtocol("pb.SyncConfigDataToAll", controller.SyncConfigDataToAll)
	grpc.RegisterProtocol("pb.PlayerLoginToGame", controller.PlayerLoginToGame)

}
