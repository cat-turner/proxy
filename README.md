# proxy

[![Go Report Card](https://goreportcard.com/badge/github.com/cat-turner/proxy)](https://goreportcard.com/report/github.com/cat-turner/proxy)

Proxy is a package to build a web server with a local cache and an external cache, like redis. The goal is to be able to get and store values, fast, using a HTTP or through a client that is RESP-compliant. This means you can proxy it up through the web or cli! Stash that data! 🐿️

## Installation

Start up service and run tests:

```bash
make test
```

Build proxy and redis contatiners and run, without running tests:

```bash
docker-compose up --build
```

If you have the redis server running in the background, and you do not have the proxy running in the background:

```bash
make build-proxy
./bin/proxy
```

Try out the RESP mode of the proxy:

```bash
make test-resp
```

Test stashing a value and key:

```bash
curl -X PUT -d "cool" localhost:8080/roxi
```

Note that the item after the route is the key you want to assign.

To get the value, simply use the key as the route.

```bash
curl localhost:8080/roxi
```

## High-level architecture overview

This module has two main components:

- proxy: A go app that has its own in-memory cache. Interacts with external cache for synchronization among other proxy instances.

- external cache: At the moment this is Redis. A cache that is run in its own docker contatiner and is also
  an in-memory data store that is run seperately from the proxy.

The build is managed with Makefile and docker-compose.yml. The app is written in go.

## What the code does

The code in this module is organized so that the main entry is clearly seperated from `proxy` package.

main.go: the entry point of the app. When configured for HTTP (APP_MODE="" or "1") it will run a http server that accepts GET and PUT requests as GET and PUT actions on the local and external cache. When configured for RESP mode (APP_MODE="2") it will accept inputs after the binary is run. This client will accept GET inputs from the inputs when it is run in this mode. This layer also configures the app to suport Sequential concurrent processing ("PROXY_CLIENT_LIMIT"=1) or Parallel concurrent processing ("PROXY_CLIENT_LIMIT"!=1).

proxy:

- config: parses variables from the environment and creates a struct that holds values that control how the app functions

- proxy: a module that has a proxy that supports Cached GET, Global expiry, LRU eviction, a fixed key capacity, and GET/PUT actions supported through HTTP

- redis: a module that interacts with a redis client to interact with a running instance of redis

- middleware: restricts number of concurrent http requests to process using buffered go channels and go routines

- cache: an interface used by the proxy. Any external cache that follows this interface can be used by the proxy to store values in an external cache.

## Algorithmic complexity of the cache operations

### Get

The proxy cache uses a map of structs to store and retireve values. Since changes to a map are not atomic, this code uses Mutex.Locks to support concurrent processing, safely. The underlying structure of maps appears to be a [hash table](http://groups.google.com/group/golang-nuts/browse_thread/thread/9286f3bc294e7ca7), so the complexity should be constant O(1). This estimate is shaky at best since there is locking implemented the Put AND Get algorithm so this may be subject to other factors such as number of clients we are concurrently processing their data and the limit set by the configuration (PROXY_CLIENT_LIMIT).

### Put

Values are stored in a map, so lookups are expected to be O(1). This mainly applies while the map is smaller that the limit set by the configuration value (CACHE_KEY_CAPACITY); it them becomes O(n) because when the size limit is reached, the LRU eviction algorithm kicks in. Other things like gargage collection and like Get, processing concurrent requests with locking, adds variabilty.

### LRU eviction

The algorithm used for LRU eviction is a variation on Least Frequently Used. The oldest element is the Less Recently Used (LRU) element. The last used timestamp is updated when an element is put into the cache or an element is retrieved from the cache with a get call. The algorithmic complexity is O(n).

### Global expiry

When the app is configured to have a global TTL ("CACHE_TTL") the proxy starts a go routine that iterates through every key, checks the time it was created, and deletes it from the map. The algorithmic complexity is O(n).

## How long you spent on each part of the project

- Planning/Research: 2
- App implementation: 8
- Writing unit tests and debugging: 5

Total: 15 hours

### Requirements not met

All requirements appear to be met, if the configurations are set correctly. This includes bonus items. If in doubt the makefile and commands can be used. By default the processing supports parallel concurrent processing, and the app needs to be configured to show sequential processing.

I am not 100% sure I am implemented the RESP portion in a way that would support all clients. This is limited to clients that care able to call on the binary using a subprocess that is seperate from that client's main application thread.
