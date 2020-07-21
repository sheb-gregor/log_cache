package metrics

import (
	"sync"
	"sync/atomic"
)

type counter struct{ val uint64 }

func (c *counter) Increment() {
	atomic.AddUint64(&c.val, 1)
}

func (c *counter) Val() uint64 {
	return atomic.LoadUint64(&c.val)
}

type UniqueIndex struct {
	mutex sync.RWMutex
	index map[string]*counter
}

func NewUniqueIndex() *UniqueIndex {
	m := UniqueIndex{}
	m.mutex = sync.RWMutex{}
	m.index = make(map[string]*counter)

	return &m
}

func (m *UniqueIndex) Add(key string) {
	m.mutex.Lock()
	if _, ok := m.index[key]; !ok {
		m.index[key] = &counter{}
	}

	m.index[key].Increment()
	m.mutex.Unlock()
}

func (m *UniqueIndex) Value(key string) uint64 {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if _, ok := m.index[key]; !ok {
		return 0
	}

	return m.index[key].Val()
}
