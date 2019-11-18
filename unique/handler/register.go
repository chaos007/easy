package handler

import "github.com/chaos007/easycome/session/grpc"

// ProtocolInit Init
func ProtocolInit() {
	grpc.RegisterProtocol("pb.PlayerLoginToUnique", PlayerLoginToUnique)
	grpc.RegisterProtocol("pb.UpJoinRoomToUnique", UpJoinRoomToUnique)

}
