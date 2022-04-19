start:
	go run main.go

test:
	go clean --testcache
	go test ./...

swag:
	swag init --parseDependency --parseInternal