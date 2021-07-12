server:
	@go run server/server.go

client:
	@go run client/main.go
	
test: 
	@go test -cover ./server

install-front:
	@cd front
	@npm i

run-front:
	@cd front
	@npm start