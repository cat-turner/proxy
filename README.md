# redis-proxy
an HTTP web service which allows the ability to add additional features on top of Redis


## Usage

GET

PUT
Put is used to create a resource or overwrite it



Assumptions

Parallel concurrent processing
The http server implementation creates a new goroutine for each incoming request. Therefore it by default handles requests concurrently.
To cap the concurrent number of requests set the limit MAX_CONNECTIONS to your desired limit. I made the assumption that when the number of requests
exceed the limit any requests made after that limit will fail. This library controls the limit of concurrent connections using a buffered channel. The capacity
of the channel buffer limits the number of simultaneous requests to process.
