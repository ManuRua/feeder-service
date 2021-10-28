package counter

import (
	"sync"
	"sync/atomic"
)

// Counter is safe to use concurrently.
type Counter struct {
	mu    sync.Mutex
	total uint64
}

// Inc increments the counter for the given key.
func (c *Counter) Inc(key string) {
	c.mu.Lock()

	atomic.AddUint64(&c.total, 1)

	c.mu.Unlock()
}

// Value returns the current value of the counter.
func (c *Counter) Value(key string) uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	return atomic.LoadUint64(&c.total)
}
