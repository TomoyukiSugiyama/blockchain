# blockchain


Generate gRPC code 
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/client.proto proto/node.proto
```

## server
```bash
# root server
go run cli/main.go -mode server -serverType root -appAddress 127.0.0.1:8080 -nodeAddress 127.0.0.1:9090
# client server
go run cli/main.go -mode server -serverType client -appAddress 127.0.0.1:8081 -nodeAddress 127.0.0.1:9091 -targetRootAddress 127.0.0.1:9090
```

### application
```bash
# request to root node
go run cli/main.go -mode app -targetNodeAddress 127.0.0.1:8080
# request to client node
go run cli/main.go -mode app -targetNodeAddress 127.0.0.1:8081
```