package client

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chaos007/easy/data/pb"
	"github.com/chaos007/easy/testscript/socket"
	"github.com/chaos007/easycome/session/web"
)

func TestLogin(t *testing.T) {
	c1 := GetClient("")
	if c1 == nil {
		return
	}
	c2 := GetClient("")
	if c2 == nil {
		return
	}

	c1.Player.Wait()
	c2.Player.Wait()
}

func GetClient(str string) *Client {
	client := NewClient()
	p := NewClientGetIPPort(str)
	if p == nil {
		return nil
	}
	client.SetAgentHost("ws://" + p.Info.AgentIP + ":" + p.Info.AgentPort + "/websocket")
	die := make(chan bool)
	client.Socket.SetServerType(socket.ConnectServerWebSocket)
	client.Do(die, client.Player)
	client.Player.ID = p.UserID
	client.Player.SetSession(client.Session)
	client.Socket.Send(&pb.PlayerLoginToUnique{
		UserID: p.UserID,
		Cookie: p.Cookie,
	})
	return client
}

// NewClientGetIPPort ip
func NewClientGetIPPort(str string) *pb.WebDownRegister {
	p := &pb.WebUpRegister{
		MachineCode: str,
	}
	j, _ := json.Marshal(p)
	m := map[string][]string{
		"data": []string{string(j)},
		"syn":  []string{web.Getmd5(string(j), web.Key)},
	}

	res, err := web.HTTPPost("http://192.168.10.89:8100/register", m)
	if err != nil {
		fmt.Println("-----------:", err)
		return nil
	}
	fmt.Println("TestLogin:", res)
	resJ := &web.JSONRet{}
	json.Unmarshal([]byte(res), resJ)

	b, _ := json.Marshal(resJ.Data)
	wb := &pb.WebDownRegister{}
	json.Unmarshal(b, wb)
	fmt.Println("TestLogin:", wb.Info.AgentIP)
	fmt.Println("TestLogin:", wb.Info.AgentPort)
	return wb
}

// GetIPPort ip
// func GetIPPort() *pb.WebDownGetServerInfo {
// 	p := &pb.WebUpGetServerInfo{
// 		UserName: "12345678",
// 		Password: "112244",
// 	}
// 	j, _ := json.Marshal(p)
// 	m := map[string][]string{
// 		"data": []string{string(j)},
// 		"syn":  []string{web.Getmd5(string(j))},
// 	}
// 	res, err := web.HTTPPost("http://192.168.0.104:8100/login", m)
// 	if err != nil {
// 		fmt.Println("-----------:", err)
// 		return nil
// 	}
// 	fmt.Println("TestLogin:", res)
// 	resJ := &web.JSONRet{}
// 	json.Unmarshal([]byte(res), resJ)

// 	b, _ := json.Marshal(resJ.Data)
// 	wb := &pb.WebDownGetServerInfo{}
// 	json.Unmarshal(b, wb)
// 	fmt.Println("TestLogin:", wb.Info.AgentIP)
// 	fmt.Println("TestLogin:", wb.Info.AgentPort)
// 	return wb
// }
