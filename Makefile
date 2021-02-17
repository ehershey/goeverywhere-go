goe: *.go
	go build

release: goe test
	GOOS=linux GOARCH=amd64 go build -o goe.linux

test: *.go
	go test -v

