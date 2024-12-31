# blockchain


Generate gRPC code 
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/blockchain.proto
```

## server
```
go run cli/main.go
```

### client
```
go run client/client.go
```