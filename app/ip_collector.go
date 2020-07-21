package app

import (
	"log_cache/metrics"

	"github.com/lancer-kit/uwe/v2"
	"github.com/sirupsen/logrus"
)

type Collector struct {
	log   *logrus.Entry
	ipBus <-chan string
	index *metrics.UniqueIndex
}

func NewCollector(ipBus <-chan string, log *logrus.Entry) *Collector {
	return &Collector{
		log:   log,
		ipBus: ipBus,
		index: metrics.NewUniqueIndex(),
	}
}

func (c *Collector) Init() error { return nil }

func (c *Collector) Run(ctx uwe.Context) error {
	c.log.Info("start ip collector daemon")

	for {
		select {
		case <-ctx.Done():
			c.index.Close()
			return nil
		case ip := <-c.ipBus:
			c.log.WithField("ip_address", ip).Debug("got new IP address")
			c.index.Add(ip)
		}
	}
}
