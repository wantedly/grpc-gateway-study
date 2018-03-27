# grpc-gateway-study
sample code for grpc-gateway study #grpc_gateway_wt

[Go + grpc-gateway でつくる JSON API 速習会 @ Wantedly](https://wantedly.connpass.com/event/82378)

## Sample

:point_right: See [develop branch](https://github.com/wantedly/grpc-gateway-study/tree/develop)

## Prerequirements

Golang (^1.9.0), dep (> 3.0.0), protobuf (> 3.0.0)

```
$ brew install protobuf
$ brew install izumin5210/tools/grapi
```

If you would like to run samples in Docker, you can use [creasty/rid](https://github.com/creasty/rid).

```
$ brew install creasty/tools/rid
```


## Links

- [Slide (in Japanese)](https://speakerdeck.com/izumin5210/create-json-api-with-go-plus-grpc-gateway)
- [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- [grapi](https://github.com/izumin5210/grapi)
  - Code generator
  - protoc wrapper
  - gRPC + grpc-gateway server implementation
