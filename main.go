package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"sync"
	"time"
)

const maxConnections = 1
const maxKeys = 1
const keyTimeout = time.Duration(100)

type ValueStore struct {
	LastRead   time.Time
	Value      string
	ExpiryTime time.Time
}

// A cache used by the proxy that is safe to use concurrently
type ProxyCache struct {
	Data map[string]ValueStore
	Mux  sync.Mutex

	// MaxKeys optionally limits the total number of keys stored in
	// the Data map at any time
	//
	// Zero means no limit
	MaxKeys int

	// KeyTimeout is the maximum amount of time the key
	// will remain in data until it is expired in Seconds
	// Zero means no limit
	KeyTimeout time.Duration
}

// Put ...
func (c *ProxyCache) Put(key string, value string) {

	if c.MaxKeys != 0 && len(c.Data) == c.MaxKeys {
		go c.PurgeKey()
	}

	c.Mux.Lock()
	defer c.Mux.Unlock()

	c.Data[key] = ValueStore{
		Value:      value,
		LastRead:   time.Now(),
		ExpiryTime: time.Now().Add(time.Second * c.KeyTimeout),
	}
}

// Get ...
func (c *ProxyCache) Get(key string) *string {
	c.Mux.Lock()

	defer c.Mux.Unlock()
	value, ok := c.Data[key]

	if ok {
		value.LastRead = time.Now()
		c.Data[key] = value
		return &value.Value
	}

	return nil
}

// PurgeKey ...
func (c *ProxyCache) PurgeKey() {
	lastKey := ""
	lastRead := time.Now()
	for k := range c.Data {
		v, ok := c.Data[k]
		if ok && v.LastRead.Before(lastRead) {
			lastKey = k
			lastRead = v.LastRead
		}
	}

	// remove the key that was accessed a longest time
	c.Mux.Lock()
	delete(c.Data, lastKey)
	c.Mux.Unlock()
}

func (c *ProxyCache) payloadHandler(w http.ResponseWriter, r *http.Request) {
	key := path.Base(r.URL.String())

	if key == "/" {
		w.Write([]byte("BAD"))
		return
	}
	switch r.Method {
	case http.MethodGet:
		value := c.Get(key)
		if value != nil {
			w.Write([]byte(*value))
		}
	case http.MethodPut:
		value, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Fatal(err)
			w.Write([]byte("BAD"))
			return
		}
		c.Put(key, string(value))
		w.Write([]byte("OK"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// limitNumClients is HTTP handling middleware that ensures no more than
// maxClients requests are passed concurrently to the given handler f.
func limitNumClients(f http.HandlerFunc, maxClients int) http.HandlerFunc {
	// buffered channel that will permit maxClients
	sema := make(chan bool, maxClients)

	return func(w http.ResponseWriter, req *http.Request) {
		sema <- true              // send value to the channel
		defer func() { <-sema }() // done; recieve value from channel but do this after func executed
		f(w, req)
	}
}

func main() {

	mux := http.NewServeMux()
	pc := ProxyCache{
		Data:       make(map[string]ValueStore),
		MaxKeys:    maxKeys,
		KeyTimeout: keyTimeout,
	}
	// limit to maxConnections for this handler
	mux.HandleFunc("/", limitNumClients(pc.payloadHandler, maxConnections))
	http.ListenAndServe(":3000", mux)
}
