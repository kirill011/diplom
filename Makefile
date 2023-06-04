pyt_gen:
	python -m grpc_tools.protoc -I=api/proto --python_out=. --pyi_out=. --grpc_python_out=. api/proto/externalApi.proto
go_gen:
	protoc -I=. --go_out=. --go-grpc_out=. --go-grpc_opt=paths=import  api/proto/externalApi.proto
serv_run:
	go run cmd/main.go
cli_run:
	go run cmd/client/client.go