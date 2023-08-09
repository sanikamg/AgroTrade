hello:
	echo "hello"

## Generate swagger docs
swagger: 
	swag init -g pkg/api/server.go -o ./cmd/api/docs

wire: 
	cd pkg/di && wire

## run command
run:
	go run ./cmd/api/main.go

golint:
	golint ./...

mockgen:
	mockgen -source=pkg/repository/interface/user.go -destination=pkg/mock/userRepoMock/userRepoMock.go -package=mock

usecasemockgen:
	mockgen -source=pkg/usecase/interface/user.go -destination=pkg/mock/userUsecaseMock/userUsecaseMock.go -package=mock

test:
	go test 

build:
	go build ./cmd/api/

##google-chrome --user-data-dir="~/chrome-dev-session" --disable-web-security
