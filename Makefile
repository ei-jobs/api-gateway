build:
	@go build -o bin/ei-jobs cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ei-jobs

seed: build
	@./bin/ei-jobs -seed
