package match

import "sync"

// RoomList 加入房间
type RoomList struct {
	PlayerList []string
	Lock       *sync.RWMutex
}

var list = &RoomList{
	PlayerList: []string{},
	Lock:       new(sync.RWMutex),
}

// JoinRoom 加入房间
func JoinRoom(name string) []string {
	list.Lock.Lock()
	defer list.Lock.Unlock()
	if len(list.PlayerList) >= 1 {
		other := list.PlayerList[0]
		list.PlayerList = list.PlayerList[1:]
		return []string{name, other}
	}
	list.PlayerList = append(list.PlayerList, name)
	return nil
}
