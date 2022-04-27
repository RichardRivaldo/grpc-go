# grpc-go

Learning gRPC with Go

### About

This is a self-exploration project regarding `gRPC` with `Protocol Buffers` in Go. I am following tutorials and videos from many sites, mainly from YouTube and Udemy (`gRPC [Golang] Master Class: Build Modern API & Microservices`). Kindly gives critics about the code or techniques used in this repository, and also suggestions about further explorations or projects (especially Backend and System Architecture, in Go or any languages) if you ever stumbled on this project. Really appreciate it! :D

### Generating Codes with Protoc

`Protoc` is a compiler specific for compiling and auto-generate the code needed to translate .proto files into language-specific codes. To run the compiler for Go, use the template below.

`protoc -I${PROJECT}/proto --go_opt=module=${YOUR_MODULE} --go_out=. ${PROJECT}/proto/*.proto --go-grpc_opt=module=${YOUR_MODULE} --go-grpc_out=. ${PROJECT}/proto/*.proto`.

Example usage for this project below.

`protoc -Idummy/proto --go_opt=module=github.com/RichardRivaldo/grpc-go --go_out=. dummy/proto/*.proto --go-grpc_opt=module=github.com/RichardRivaldo/grpc-go --go-grpc_out=. dummy/proto/*.proto`
