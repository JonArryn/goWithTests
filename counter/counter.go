package counter

import "sync"

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Value() int {
	return c.count
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
}
