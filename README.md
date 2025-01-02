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
go run cli/main.go -mode server -type master
# client server
go run cli/main.go -mode server -type client
```

### client
```bash
go run cli/main.go -mode client -address 127.0.0.1:8080
```