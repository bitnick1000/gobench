package sync

import (
	"sync"
	"testing"
)

type Plus struct {
	lock *sync.RWMutex
	i    int
}
type Plus2 struct {
	lock *sync.Mutex
	i    int
}

func Benchmark_Mutex_Lock(b *testing.B) {
	plus := Plus2{new(sync.Mutex), 0}
	for i := 0; i < b.N; i++ {
		plus.lock.Lock()
		plus.i++
		plus.lock.Unlock()
	}
}

func Benchmark_RWMutex_Lock(b *testing.B) {
	plus := Plus{new(sync.RWMutex), 0}
	for i := 0; i < b.N; i++ {
		plus.lock.Lock()
		plus.i++
		plus.lock.Unlock()
	}
}

func Benchmark_RWMutex_RLock(b *testing.B) {
	plus := Plus{new(sync.RWMutex), 0}
	for i := 0; i < b.N; i++ {
		plus.lock.RLock()
		plus.i++
		plus.lock.RUnlock()
	}
}

func Benchmark_nolock(b *testing.B) {
	plus := Plus{new(sync.RWMutex), 0}
	for i := 0; i < b.N; i++ {
		plus.i++
	}
}
