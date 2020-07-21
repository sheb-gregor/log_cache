package metrics

import (
	"sync"
	"sync/atomic"

	"github.com/prometheus/client_golang/prometheus"
)

const CounterUniqueIP string = "unique_ip_addresses"

// counter is trade-safe counter.
type counter struct{ val uint64 }

// Increment adds 1 to value.
func (c *counter) Increment() {
	atomic.AddUint64(&c.val, 1)
}

// Val returns current state of counter.
func (c *counter) Val() uint64 {
	return atomic.LoadUint64(&c.val)
}

// UniqueIndex is trade-safe index/counter.
type UniqueIndex struct {
	mutex   sync.RWMutex
	index   map[string]*counter
	counter prometheus.Counter
}

// NewUniqueIndex initializes UniqueIndex properly.
func NewUniqueIndex() *UniqueIndex {
	m := UniqueIndex{}
	m.mutex = sync.RWMutex{}
	m.index = make(map[string]*counter)

	m.counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name:        CounterUniqueIP,
			ConstLabels: map[string]string{},
		})

	prometheus.MustRegister(m.counter)
	return &m
}

// Add increments counter for key, if key not present - initializes and increments unique metric.
func (m *UniqueIndex) Add(key string) {
	m.mutex.Lock()
	if _, ok := m.index[key]; !ok {
		m.index[key] = &counter{}
		m.counter.Inc()
	}

	m.index[key].Increment()
	m.mutex.Unlock()
}

// Value returns counter for key, if key not present - returns 0.
func (m *UniqueIndex) Value(key string) uint64 {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if _, ok := m.index[key]; !ok {
		return 0
	}

	return m.index[key].Val()
}

func (m *UniqueIndex) Close() {
	prometheus.Unregister(m.counter)
}
