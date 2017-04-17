test: lint
	@go test -v ./...

lint: golint vet

golint:
	@golint ./...

vet:
	@go vet ./...
