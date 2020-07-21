package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type MKey string

type Collector struct {
	counters map[MKey]prometheus.Counter
}

const CounterUniqueIP MKey = "unique_ip_addresses"

var (
	enabled   bool      // nolint:gochecknoglobals
	collector Collector // nolint:gochecknoglobals
)

func Init() {
	enabled = true
	collector = Collector{
		counters: map[MKey]prometheus.Counter{},
	}

	for _, name := range []MKey{CounterUniqueIP} {
		counter := prometheus.NewCounter(
			prometheus.CounterOpts{
				Name:        string(name),
				ConstLabels: map[string]string{},
			})

		prometheus.MustRegister(counter)
		collector.counters[name] = counter
	}
}

// Inc increments the Gauge by 1. Use Add to increment it by arbitrary
// values.
func Inc(name MKey) {
	if !enabled {
		return
	}
	collector.counters[name].Inc()
}

// Add adds the given value to the Gauge. (The value can be negative,
// resulting in a decrease of the Gauge.)
func Add(name MKey, val float64) {
	if !enabled {
		return
	}
	collector.counters[name].Add(val)
}
