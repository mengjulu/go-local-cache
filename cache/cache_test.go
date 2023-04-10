package localcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	data = map[string]interface{}{
		"str":   "test",
		"num":   123,
		"float": 1.23,
		"bool":  false,
		"arr":   []string{"a", "b", "c"},
		"obj":   map[string]string{"a": "b", "c": "d"},
		"struct": struct {
			A string
		}{"test"},
	}
	newData = map[string]interface{}{
		"str": "new test",
		"num": 789,
	}
	c = New()
)

func TestSetCache(t *testing.T) {
	ok := c.Set("data", data)
	assert.True(t, ok)
}

func TestGetCache(t *testing.T) {
	c.Set("data", data)
	val, ok := c.Get("data")
	assert.Equal(t, data, val)
	assert.True(t, ok)
}

func TestCacheOverwrite(t *testing.T) {
	c.Set("data", newData)
	val, ok := c.Get("data")
	assert.Equal(t, newData, val)
	assert.True(t, ok)
}

func TestCacheExpire(t *testing.T) {
	c.Set("data", data)
	time.Sleep(exp + 1*time.Second)
	val, ok := c.Get("data")
	assert.Nil(t, val)
	assert.False(t, ok)
}
