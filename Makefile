hello:
	echo "hello"

## Generate swagger docs
swagger: 
	swag init -g pkg/api/server.go -o ./cmd/api/docs

wire: 
	cd pkg/di && wire

run:
	go run ./cmd/api/main.go

golint:
	golint ./...

##google-chrome --user-data-dir="~/chrome-dev-session" --disable-web-security
