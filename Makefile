
test-multi:
	for i in `seq 1 1000`; \
	do\
		curl  localhost:3000/cache;\
	done;\