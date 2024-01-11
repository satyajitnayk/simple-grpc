# simple-grpc

A simple GRPC server &amp; client using go lang

## Use command to generate .pb.go files

```cmd
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```
