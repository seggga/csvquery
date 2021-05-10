COMMIT_VER := $(shell git rev-list -1 HEAD)

check:
	${HOME}/go/bin/golangci-lint run

.PHONY: build

build:
	mkdir -p ./bin
	go build -o ./bin/csvquery -ldflags "-X main.gitCommit=$(COMMIT_VER)" ./main.go ./scan_csv.go ./init_logger.go ./binary_data.go

test:
	go test -cover ./...
