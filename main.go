package main

import (
	"net/http"
	"fmt"
)

const maxConnections = 1

func payloadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("OK"))
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
	// Counting semaphore using a buffered channel
	sema := make(chan struct{}, maxClients)

	return func(w http.ResponseWriter, req *http.Request) {
	  sema <- struct{}{}
	  defer func() { <-sema }()
	  f(w, req)
	}
  }

func main() {
	mux := http.NewServeMux()
	// limit to maxConnections for this handler
    mux.HandleFunc("/cache", limitNumClients(payloadHandler,maxConnections))
	http.ListenAndServe(":3000", mux)
}