jcconv: *.go */*.go
	go build -o jcconv

linux: *.go */*.go
	env GOOS=linux GOARCH=amd64 go build -o jcconv-linux

dev: *.go */*.go
	go build -race -o jcconv

dep:
	GOOS=windows go get -u github.com/spf13/cobra

gen: box/resources
	go generate ./...

lint:
	golint src/...
	golangci-lint run

all: jcconv jcconv-linux

clean:
	rm jcconv jcconv-linux
