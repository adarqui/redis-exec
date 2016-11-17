build: deps
	go build

deps:
	go get gopkg.in/redis.v4

test:
	go test

clean:
	rm -f redis-exec
