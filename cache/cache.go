// this package offers set and get function for localcache. The data will expire after 30 seconds and can be overwritten.
package localcache

import (
	"sync"
	"time"
)

var (
	exp = 30 * time.Second
)

// localcache is the implementation of Cache interface with stored data, expiration time and mutex.
type localcache struct {
	data map[string]interface{}
	exp  time.Duration
	mut  sync.RWMutex
}

// Get retrieves data from the cache by key.
func (c *localcache) Get(key string) (interface{}, bool) {
	c.mut.RLock()
	defer c.mut.RUnlock()

	val, ok := c.data[key]
	return val, ok
}

// Set stores data with key and will be deleted after 30 seconds automatically.
func (c *localcache) Set(key string, val interface{}) bool {
	c.mut.Lock()
	defer c.mut.Unlock()

	time.AfterFunc(exp, func() {
		c.mut.Lock()
		defer c.mut.Unlock()
		delete(c.data, key)
	})

	c.data[key] = val
	return true
}

func New() Cache {
	return &localcache{
		data: make(map[string]interface{}),
		exp:  exp,
	}
}
