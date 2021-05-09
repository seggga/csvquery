test:
	go test -cover ./...

run:
	go run .

check:
	golangci-lint run
