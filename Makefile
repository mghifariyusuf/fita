vendor:
	@go mod vendor

test:
	@go test ./test/... -mod=vendor -coverpkg=./... -failfast -short -cover -coverprofile=coverage.out
	@go tool cover -func=coverage.out | tail -n 1 | awk '{ print $3 }'

build:
	@go build -mod=vendor -v -o fita cli/main.go

test-and-build:
	@go test ./test/... -mod=vendor -coverpkg=./... -failfast -short -cover -coverprofile=coverage.out
	@go tool cover -func=coverage.out | tail -n 1 | awk '{ print $3 }'
	@go build -mod=vendor -v -o fita cli/main.go