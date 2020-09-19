package main

import (
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })
    http.ListenAndServe(":3000", mux)
}