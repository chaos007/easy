package player

import (
	"fmt"
	"sync"

	"github.com/chaos007/easy/testscript/types"
)

// PlayersManager Players map[int64]*Player
var PlayersManager = NewPlayers()

// Players player
type Players struct {
	data map[string]*Player
	lock sync.RWMutex
}

// NewPlayers NewPlayers
func NewPlayers() *Players {
	return &Players{
		data: make(map[string]*Player),
	}
}

// Get Get
func (p *Players) Get(id string) *Player {
	p.lock.Lock()
	defer p.lock.Unlock()
	if a, ok := p.data[id]; ok {
		return a
	}
	return nil
}

// Add Add
func (p *Players) Add(a *Player) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.data[a.ID] = a
}

// Delete Delete
func (p *Players) Delete(id string) {
	p.lock.Lock()
	defer p.lock.Unlock()
	delete(p.data, id)
}

// Player Player
type Player struct {
	ID               string
	NickName         string
	CurrentGameState int32
	Session          *types.Session
}

// SetSession SetSession
func (p *Player) SetSession(sess *types.Session) {
	p.Session = sess
}

// Done Done
func (p *Player) Done(method interface{}) {
	fmt.Printf("%s done\n", method)
	p.Session.Step <- method
}

// Wait Wait
func (p *Player) Wait() interface{} {
	method := <-p.Session.Step
	fmt.Printf("wait for after %s\n", method)
	return method
}

// GetPlayer GetPlayer
func GetPlayer(userID string) *Player {
	if p := PlayersManager.Get(userID); p != nil {
		return p
	}
	p := NewPlayer()
	p.ID = userID
	PlayersManager.Add(p)
	return p
}

// NewPlayer NewPlayer
func NewPlayer() *Player {
	return &Player{
		NickName: "",
		// SaveList: make([]pb.PBSaveListItem, 0),
	}
}
