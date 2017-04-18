test: lint
	@go test -v ./...

lint: golint vet

golint:
	@golint ./...

vet:
	@go vet ./...

vmrun: cmd/vmrun/main.go
	@go run cmd/vmrun/main.go
