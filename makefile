jcc: *.go */*.go
	go build -o jcc

linux: *.go */*.go
	env GOOS=linux GOARCH=amd64 go build -o jcc-linux

dev: *.go */*.go
	go build -race -o jcc

gen: box/resources
	go generate ./...

lint:
	golint src/...
	golangci-lint run

all: jcc jcc-linux

clean:
	rm jcc jcc-linux
