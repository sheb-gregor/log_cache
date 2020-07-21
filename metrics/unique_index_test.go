package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueIndex_Value(t *testing.T) {
	index := NewUniqueIndex()
	ip1 := "0.0.0.1"
	ip2 := "0.0.0.2"

	assert.Equal(t, uint64(0), index.Value(ip1))
	assert.Equal(t, uint64(0), index.Value(ip2))

	index.Add(ip1)
	assert.Equal(t, uint64(1), index.Value(ip1))
	assert.Equal(t, uint64(0), index.Value(ip2))

	index.Close()
}

func TestUniqueIndex_Add(t *testing.T) {
	index := NewUniqueIndex()
	ip1 := "0.0.0.1"
	ip2 := "0.0.0.2"
	ip3 := "0.0.0.3"

	index.Add(ip1)
	assert.Equal(t, uint64(1), index.Value(ip1))
	assert.Equal(t, uint64(0), index.Value(ip2))
	assert.Equal(t, uint64(0), index.Value(ip3))

	index.Add(ip1)
	index.Add(ip2)

	assert.Equal(t, uint64(2), index.Value(ip1))
	assert.Equal(t, uint64(1), index.Value(ip2))
	assert.Equal(t, uint64(0), index.Value(ip3))

	index.Add(ip3)
	assert.Equal(t, uint64(2), index.Value(ip1))
	assert.Equal(t, uint64(1), index.Value(ip2))
	assert.Equal(t, uint64(1), index.Value(ip3))

	index.Close()
}
