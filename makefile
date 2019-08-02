jcc: src/*.go
	go build -o jcc ./src

linux: src/*.go
	env GOOS=linux GOARCH=amd64 go build -o jcc-linux ./src

dev: src/*.go
	go build -race -o jcc ./src

gen: box/resources
	go generate ./...

lint:
	golint src/...
	golangci-lint run

all: jcc jcc-linux

clean:
	rm jcc jcc-linux
