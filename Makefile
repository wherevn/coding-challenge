
dep:
	@go mod download
	@go mod vendor

build:
	@go build ./...

test:
	@go test -cover ./...
