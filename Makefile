.PHONY: run test grpcurl-echo

run:
	go run cmd/server/main.go

test:
	go test -v ./...

grpcurl-echo:
	grpcurl -plaintext -d '{"message":"hoge"}' localhost:8080 echo.v1.EchoService/Echo
