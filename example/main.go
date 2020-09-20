package main

import (
	"net/http"
	"time"
)

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

	maxConnections := 1
	maxKeys := 1
	keyTimeout := time.Duration(1000)

	// create a new instance of the proxy cache
	pc := proxy.NewProxyCache(&maxKeys, &keyTimeout)

	// limit to maxConnections for this handler
	mux.HandleFunc("/", limitNumClients(pc.payloadHandler, maxConnections))
	http.ListenAndServe(":3000", mux)
}
