package concurrentMap

import (
	"context"
	"sync"
	"time"
)

type ConcurrentMap struct {
	mux      sync.Mutex
	data     map[any]any
	infoChan map[any]chan struct{}
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		data:     make(map[any]any),
		infoChan: make(map[any]chan struct{}),
	}
}

func (c *ConcurrentMap) Put(key, val any) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.data[key] = val
	if ch, ok := c.infoChan[key]; ok {
		close(ch)
		delete(c.infoChan, key)
	}
}

func (c *ConcurrentMap) Get(key any, timeOut time.Duration) (any, bool) {
	c.mux.Lock()
	if v, ok := c.data[key]; ok {
		c.mux.Unlock()
		return v, true
	}

	var ch chan struct{}
	if _, ok := c.infoChan[key]; !ok {
		c.infoChan[key] = make(chan struct{})
		ch = c.infoChan[key]
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	c.mux.Unlock()

	select {
	case <-ctx.Done():
		return nil, false
	case <-ch:
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.data[key], true
}
