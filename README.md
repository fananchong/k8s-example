# k8s-example
使用kubernets部署典型服务架构的例子

## 例子1

`服务A` 通过k8s服务发现，获取 `服务B` 列表，并发送1个地址给 客户端

`服务B`一台机器只允许开1个容器

参见 [example1/README.md](example1/README.md)


## 例子2

`服务A` 通过k8s服务发现，获取 `服务B` 列表，并发送1个地址给 客户端  

`服务B`一台机器不限制仅开1个容器

参见 [example2/README.md](example2/README.md)


## 例子3

`服务A` 通过k8s服务发现，获取 `服务B` 列表，并与 `服务B` 互连

参见 [example3/README.md](example3/README.md)
