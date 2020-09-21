
test-multi-get:
	for i in `seq 1 1000`; \
	do\
		curl  localhost:8080/cache;\
	done;\


test-multi-put:
	for i in `seq 1 1000`; \
	do\
		curl -d "dataNew" -X PUT localhost:3000/cache;\
	done;\


test-multi-get2:
	for i in `seq 1 1000`; \
	do\
		curl  localhost:8080/cache2;\
	done;\


test-multi-put2:
	for i in `seq 1 1000`; \
	do\
		curl -d "dataOtherNew" -X PUT localhost:3000/cache2;\
	done;\

test-multi-get3:
	for i in `seq 1 1000`; \
	do\
		curl  localhost:8080/cache3;\
	done;\


test-multi-put3:
	for i in `seq 1 1000`; \
	do\
		curl -d "Roxi" -X PUT localhost:8080/cache3;\
	done;\


test: test-multi-get test-multi-put test-multi-get2 test-multi-put2 test-multi-get3 test-multi-put3


all:
	docker-compose up --build -d

build-proxy:
	mkdir -p bin
	go build -o ./bin/proxys .

build-proxy-sequental-processing:
	mkdir -p bin
	PROXY_CLIENT_LIMIT=1 go build -o ./bin/proxys .