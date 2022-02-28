run_server:
	go run pkg/server/main.go

run_stream:
	go run pkg/streamer/main.go

compile_proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/proto_db/proto_db.proto
