.PHONY: run windows mac linux

run:
	go build ./main.go
	
windows:
	GOOS=windows GOARCH=amd64 go build ./main.go

mac:
	GOOS=darwin GOARCH=amd64 go build ./main.go

linux:
	GOOS=linux GOARCH=amd64 go build ./main.go