# blockchain


Generate gRPC code 
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/client.proto proto/node.proto
```

## server
```bash
# master server
go run cli/main.go -type master
# client server
go run cli/main.go -type client
```

### client
```bash
go run client/client.go
```