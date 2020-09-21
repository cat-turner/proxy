


all:
	docker-compose up --build -d

build-proxy:
	mkdir -p bin
	go build -o ./bin/proxy .

build-proxy-sequental-processing:
	mkdir -p bin
	PROXY_CLIENT_LIMIT=1 go build -o ./bin/proxy .

test:
	# The test should test the Redis proxy in its running state (i.e. by starting the
	# artifact that would be started in production)
	docker-compose up -d
	go test ./...