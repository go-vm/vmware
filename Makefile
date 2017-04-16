lint: golint vet

golint:
	@golint ./...

vet:
	@go vet ./...
