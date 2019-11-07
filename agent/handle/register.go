package handle

import (
	"github.com/chaos007/easy/agent/handle/controller"
	"github.com/chaos007/easycome/session/grpc"
	"github.com/chaos007/easycome/session/mixkcp"
)

// Init Init
func Init() {
	grpc.RegisterProtocol("pb.TestToAgent", controller.TestToAgent)
	mixkcp.RegisterProtocol("pb.TestStruct", controller.TestStruct)
	mixkcp.RegisterProtocol("pb.UpPing", controller.UpPing)

	mixkcp.RegisterProtocol("pb.PlayerLogin", controller.PlayerLogin)
	mixkcp.RegisterProtocol("pb.LoginCheckPlayerToAgent", controller.LoginCheckPlayerToAgent)
	grpc.RegisterProtocol("pb.LoginKickPlayerToAgent", controller.LoginKickPlayerToAgent)

}
