default: test

install:
	@go install -v -x ./...

test: lint
	@go test -v ./...

lint: golint vet

golint:
	@golint ./...

vet:
	@go vet ./...

coverage:
	@./scripts/coverage.bash

vmruntest: cmd/vmruntest/main.go
	@go run cmd/vmruntest/main.go
