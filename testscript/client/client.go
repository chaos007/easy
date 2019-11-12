package client

import (
	"fmt"
	"os"
	"time"

	"github.com/chaos007/easy/data/pb"
	"github.com/chaos007/easy/testscript/player"
	"github.com/chaos007/easy/testscript/socket"
	"github.com/chaos007/easy/testscript/types"
	
)

// Client 客户端
type Client struct {
	Session *types.Session
	Socket  *socket.Socket
}

// LogFile LogFile
var LogFile *os.File

// SetAgentHost 设置头
func (c *Client) SetAgentHost(host string) {
	c.Socket.SetHost(host)
}

// Do 通信
func (c *Client) Do(die chan bool, p *player.Player) (err error) {
	if err = c.Socket.Dail(); err != nil {
		fmt.Println("--------err:", err)
		fmt.Printf("-------Dail server failed!")
	} else {
		go c.Socket.RecvThread(c.Session, die, p)
	}
	return
}

// Play Play
func (c *Client) Play(p *player.Player, battleType int32) string {
	return "Down"
}

func uppingtimer(c *Client) {
	timer2 := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-timer2.C:
			c.Socket.Send(&pb.UpPing{})
		}
	}
}

// NewClient NewClient
func NewClient() *Client {
	return &Client{
		Session: types.NewSession(),
		Socket:  socket.NewSocket(),
	}
}

// UpPlayerAction UpPlayerAction
func (c *Client) UpPlayerAction() {
	i := 0
	tick := time.Tick(1 * time.Second)
	for {
		select {
		case <-tick:
			if i%100 == 0 {
				fmt.Printf("-------------rece from tick:%d,userid:%d\n", i, c.Session.UserID)
			}
			i++
		}
	}

}
