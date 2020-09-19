package main

import (
	"net/http"
	"time"
)

const maxConnections = 1

func payloadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("OK"))
		time.Sleep(2 * time.Second)
	case "PUT":
		w.Write([]byte("OK"))
	case "POST":
		w.WriteHeader(http.StatusMethodNotAllowed)
	case "DELETE":
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
	// limit to maxConnections for this handler
	mux.HandleFunc("/cache", limitNumClients(payloadHandler, maxConnections))
	http.ListenAndServe(":3000", mux)
}
