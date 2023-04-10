package localcache

// Cache is a interface with function get and set for localcache
type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, val interface{}) bool
}
