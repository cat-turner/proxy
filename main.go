package main

import (
	"net/http"

	"github.com/cat-turner/proxy/proxy"
)

func main() {

	mux := http.NewServeMux()

	configs := proxy.NewConfig()

	// create a new instance of the proxy cache
	pc := proxy.NewProxyCache(configs)

	if configs.ProxyClientLimit != nil {
		mux.HandleFunc("/", proxy.LimitNumClients(pc.PayloadHandler, *configs.ProxyClientLimit))
	} else {
		mux.HandleFunc("/", pc.PayloadHandler)
	}

	http.ListenAndServe(configs.Port, mux)
}
