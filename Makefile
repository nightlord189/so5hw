start:
	go run main.go

test:
	go clean --testcache
	go test ./...

swag:
	swag init --parseDependency --parseInternal

build-frontend:
	cd web; npm run build;

start-frontend:
	cd web; npm start;

install-frontend:
	cd web; npm install;