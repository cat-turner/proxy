package proxy

import "net/http"

// LimitNumClients is HTTP handling middleware that ensures no more than
// maxClients requests are passed concurrently to the given handler f.
func LimitNumClients(f http.HandlerFunc, maxClients int) http.HandlerFunc {
	// buffered channel that will permit maxClients
	sema := make(chan bool, maxClients)

	return func(w http.ResponseWriter, req *http.Request) {
		sema <- true              // send value to the channel
		defer func() { <-sema }() // done; receive value from channel but do this after func executed
		f(w, req)
	}
}
