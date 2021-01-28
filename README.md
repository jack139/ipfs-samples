## IPFS API示例

### 1. 使用coreapi
   不需要ipfs daemon，可以直接读写数据

```
go run ipfs_coreapi.go
```

### 2. 使用http api
需要先启动ipfs daemon
```
nohup ipfs daemon --enable-namesys-pubsub > /tmp/ipfs.log 2>&1 &

go run ipfs_api.go
```
