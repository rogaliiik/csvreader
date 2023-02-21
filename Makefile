build:
	go build -o csvreader.exe .

test:
	go test

.DEFAULT_GOAL := build