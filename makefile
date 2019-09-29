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
	golint ./...
	golangci-lint run --skip-dirs box

all: jcconv jcconv-linux

clean:
	rm jcconv jcconv-linux

publish:
	make linux
	ssh go-pg 'sudo systemctl stop jcconv'
	scp jcconv-linux go-pg:~/jcconv/jcconv
	ssh go-pg 'sudo systemctl start jcconv'
