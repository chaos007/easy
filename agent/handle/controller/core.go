package controller

import (
	"github.com/chaos007/easy/data/pb"
	"github.com/chaos007/easycome/session/grpc"
	"github.com/chaos007/easycome/session/mixkcp"

	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
)

// TestToAgent test
func TestToAgent(sess *grpc.Session, data proto.Message) (proto.Message, error) {
	source := data.(*pb.TestToAgent)

	logrus.Errorln("----------TestToAgent:%#v", source.FieldA)

	result := &pb.TestStruct{
		FieldA: "44444",
	}

	s := mixkcp.GetUserSession(sess.UserID)
	if s != nil {
		s.SendToClientUseProto(result)
	}

	// sess.SendToSteamServer("", result)

	return nil, nil
}

// TestStruct test
func TestStruct(sess *mixkcp.Session, data proto.Message) (proto.Message, error) {
	source := data.(*pb.TestStruct)

	logrus.Errorln("----------TestStruct:%#v", source.FieldA)

	result := &pb.TestStruct{
		FieldA: "333",
	}

	// toGame := &pb.TestToGame{
	// 	FieldA: "2222",
	// }

	// sess.SendToServer(enum.ServerTypeGame, toGame)

	// sess.SendToServer(enum.ServerTypeCenter, &pb.UpGetConfigToCenter{})

	return result, nil
}

// UpPing test
func UpPing(sess *mixkcp.Session, data proto.Message) (proto.Message, error) {
	_, ok := data.(*pb.UpPing)
	if !ok {
		return nil, nil
	}

	logrus.Debugln("-----------:UpPing")
	result := &pb.DownPing{}

	return result, nil
}

// LoginKickPlayerToAgent 被通知有玩家登陆
func LoginKickPlayerToAgent(sess *grpc.Session, data proto.Message) (proto.Message, error) {
	source, ok := data.(*pb.LoginKickPlayerToAgent)
	if !ok {
		return nil, nil
	}

	s := mixkcp.GetUserSession(source.UserID)
	if s != nil {
		s.SessionClose()
	}

	return nil, nil
}

// LoginCheckPlayerToAgent LoginCheckPlayerToAgent
func LoginCheckPlayerToAgent(sess *mixkcp.Session, data proto.Message) (proto.Message, error) {
	source, ok := data.(*pb.LoginCheckPlayerToAgent)
	if !ok {
		return nil, nil
	}

	logrus.Debugln("source:", source.UserID)
	logrus.Debugln("source:", source.IsLegal)

	if !source.IsLegal { //踢掉原来的链接
		s := mixkcp.GetUserSession(source.UserID)
		if s != nil {
			s.SendToClientUseProto(&pb.DownPlayerLogin{
				IsSucceed: false,
			})
			s.SessionClose()
		}
	} else { //登陆，当前所处理的为原来mixkcp的session
		sess.SessionLogin(source.UserID)
		sess.SendToClientUseProto(&pb.DownPlayerLogin{
			IsSucceed: true,
			UserID:    source.UserID,
		})
		// sess.SendToServer(enum.ServerTypeGame, &pb.PlayerLoginToGame{
		// 	UserID: source.UserID,
		// })
	}

	return nil, nil
}
