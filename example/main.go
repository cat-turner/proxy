package main

import (
	"net/http"
	"time"

	"github.com/cat-turner/proxy"
)

func main() {

	mux := http.NewServeMux()

	maxConnections := 1
	maxKeys := 1
	keyTimeout := time.Duration(1000)

	// create a new instance of the proxy cache
	pc := proxy.NewProxyCache(&maxKeys, &keyTimeout)

	// limit to maxConnections for this handler
	mux.HandleFunc("/", proxy.limitNumClients(pc.PayloadHandler, maxConnections))
	http.ListenAndServe(":3000", mux)
}
