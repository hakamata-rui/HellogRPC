# HellogRPC

Example of gRPC for golang using protocol buffers.

https://github.com/grpc/grpc-go
https://github.com/google/protobuf

# Usage

## install

```
make install
```

To compile .proto file, download a pre-built binary.
https://github.com/google/protobuf/releases

## run

```
$ make run
go run server/main.go & 
sleep 5
go run client/main.go 
message:"hello, world" 

```