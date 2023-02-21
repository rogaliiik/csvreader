build:
	go build -o csvreader.exe ./cmd

test:
	go test ./cmd

.DEFAULT_GOAL := build