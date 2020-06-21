# 概念

- IDL: 接口定义语言
- 服务定义, 四种服务
    - Unary RPCs
    - Server streaming RPCs
    - Client streaming RPCs
    - Bidirectional streaming RPCs

# 相关命令

```shell script
# 记录一些命令参数
$ protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative helloworld/helloworld.proto

$ protoDir="../protos"
$ outDir="../languages/golang/gym"
$ protoc -I ${protoDir}/ ${protoDir}/*proto --go_out=plugins=grpc:${outDir}
```
