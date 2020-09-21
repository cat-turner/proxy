package proxy

// Cache is an interface that is not the in-memory cache used by the proxy
// also known as the external cache, like redis
type Cache interface {
	Put(key string, value string) error
	Get(key string) (*string, error)
}
