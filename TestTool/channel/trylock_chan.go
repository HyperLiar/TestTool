package main

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mu := &Mutex{make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

func (m *Mutex) Lock() {
	<-m.ch
}

func (m *Mutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

// channel边界条件会阻塞
func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}

	return false
}

func (m *Mutex) IsLocked() bool {
	return len(m.ch) == 0
}
