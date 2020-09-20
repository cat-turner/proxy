
test-multi-get:
	for i in `seq 1 1000`; \
	do\
		curl  localhost:3000/cache;\
	done;\


test-multi-put:
	for i in `seq 1 1000`; \
	do\
		curl -d "dataNew" -X PUT localhost:3000/cache;\
	done;\


test-multi-get2:
	for i in `seq 1 1000`; \
	do\
		curl  localhost:3000/cache2;\
	done;\


test-multi-put2:
	for i in `seq 1 1000`; \
	do\
		curl -d "dataNew" -X PUT localhost:3000/cache2;\
	done;\