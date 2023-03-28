.PHONY: init run grpcurl-echo

init:


run:
	go run cmd/server/main.go

grpcurl-echo:
	grpcurl -plaintext -d '{"message":"hoge"}' localhost:8080 echo.v1.EchoService/Echo
