package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type PlusMutex struct {
	lock   *sync.Mutex
	number int
}

func Test_(t *testing.T) {
	plus := &PlusMutex{new(sync.Mutex), 0}
	plus.lock.Lock()
	timeNow("1")
	go func() {
		timeNow("in func 1")
		plus.lock.Lock()
		timeNow("in func 2")
		// time.Sleep(time.Second)
		plus.number++
		plus.lock.Unlock()
	}()
	timeNow("3")
	time.Sleep(time.Millisecond * 100)
	plus.lock.Unlock()
}

func timeNow(flag string) {
	fmt.Println(flag + " " + time.Now().Format("5.000000"))
}
