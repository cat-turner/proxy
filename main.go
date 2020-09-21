package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/cat-turner/proxy/proxy"
)

func main() {

	configs := proxy.NewConfig()
	// create a new instance of the proxy cache
	pc := proxy.NewProxyCache(configs)

	if configs.Mode == "2" {
		// Special mode that allows you to interact with the proxy through the cli
		fmt.Println("RESPY>")

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			fmt.Println(input)
			keyInput := strings.Split(input, " ")
			if string(keyInput[0]) != "GET" {
				fmt.Println("Only GET is supported")

			} else {
				value, err := pc.HandleGet(string(keyInput[1]))
				if err != nil {
					fmt.Println(err)
				}
				if value == nil {
					fmt.Println("(nil)")
				} else {
					fmt.Println(*value)
				}

			}

		}

		if scanner.Err() != nil {
			fmt.Println(scanner.Err())
			return
		}

	}

	mux := http.NewServeMux()

	if configs.ProxyClientLimit != nil {
		mux.HandleFunc("/", proxy.LimitNumClients(pc.PayloadHandler, *configs.ProxyClientLimit))
	} else {
		mux.HandleFunc("/", pc.PayloadHandler)
	}

	http.ListenAndServe(configs.Port, mux)
}
