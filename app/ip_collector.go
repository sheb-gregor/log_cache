package app

import (
	"log_cache/metrics"

	"github.com/lancer-kit/uwe/v2"
)

type Collector struct {
	ipBus <-chan string
	index *metrics.UniqueIndex
}

func NewCollector(ipBus <-chan string) *Collector {
	return &Collector{
		ipBus: ipBus,
		index: metrics.NewUniqueIndex(),
	}
}

func (c *Collector) Init() error { return nil }

func (c *Collector) Run(ctx uwe.Context) error {
	for {
		select {
		case <-ctx.Done():
			c.index.Close()
			return nil
		case ip := <-c.ipBus:
			c.index.Add(ip)
		}
	}
}
