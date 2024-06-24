.PHONY: run windows mac linux

run:
	go build ./cmd/main.go
	
windows:
	GOOS=windows GOARCH=amd64 go build ./cmd/main.go

mac:
	GOOS=darwin GOARCH=amd64 go build ./cmd/main.go

linux:
	GOOS=linux GOARCH=amd64 go build ./cmd/main.go