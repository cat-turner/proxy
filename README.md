# redis-proxy
an HTTP web service which allows the ability to add additional features on top of Redis


## Usage

GET
```
curl  localhost:PORT/key
```


PUT
Put is used to create a resource or overwrite it. The key should be given in the url because that is where we would expect GET to be [successful](https://tools.ietf.org/html/rfc7231#section-4.3.4)
```
curl -d "data to PUT" -X PUT http://localhost:PORT/key
```


local cache
mutex for caches because I need to lock when writing to the map from the goroutine


Assumptions

Parallel concurrent processing
The http server implementation creates a new goroutine for each incoming request. Therefore it by default handles requests concurrently.
To cap the concurrent number of requests set the limit MAX_CONNECTIONS to your desired limit. I made the assumption that when the number of requests
exceed the limit any requests made after that limit will be blocked until the total number of requests at the time are below the limit. This library controls the limit of concurrent connections using a buffered channel in a function used as middleware. The capacity of the channel buffer limits the number of simultaneous requests to process.

