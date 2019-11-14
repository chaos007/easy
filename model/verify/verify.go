package verify

import (
	"sync"
	"time"
)

var l = &ListVerifyNumber{
	lock: new(sync.RWMutex),
	M:    map[string]*Verify{},
}

// GetVerifyMap 获得验证表
func GetVerifyMap() *ListVerifyNumber {
	return l
}

// ListVerifyNumber 验证码表
type ListVerifyNumber struct {
	lock *sync.RWMutex
	M    map[string]*Verify
}

// Add 添加验证码表
func (l *ListVerifyNumber) Add(check, number string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.clean()
	n, ok := l.M[check]
	if !ok {
		p := &Verify{
			Check:  check,
			Number: number,
			Unix:   time.Now().Add(5 * time.Minute).Unix(),
		}
		l.M[check] = p
		return
	}
	n.Number = number
	n.Unix = time.Now().Add(5 * time.Minute).Unix()
}

// Check 检查验证码表
func (l *ListVerifyNumber) Check(check, number string) bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.clean()
	if n, ok := l.M[check]; !ok {
		return false
	} else if number != n.Number {
		return false
	} else if time.Now().Unix() > n.Unix {
		return false
	}
	return true
}

func (l *ListVerifyNumber) clean() {
	for _, item := range l.M {
		if time.Now().Unix() > item.Unix {
			delete(l.M, item.Check)
		}
	}
}

// Verify 验证表
type Verify struct {
	Check  string
	Number string
	Unix   int64
}
