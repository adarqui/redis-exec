build: deps
	go build

build-static: deps
	CGO_ENABLED=0 GOARCH=amd64 go build -v -ldflags '-d -s -w' -o ./redis-exec

deps:
	go get gopkg.in/redis.v4

test:
	go test

clean:
	rm -f redis-exec
