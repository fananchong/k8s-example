## helm介绍

本例子是使用 k8s helm 来部署 redis 主从

关于 helm 如何安装，可以参考：[helm安装介绍（备忘）](https://blog.csdn.net/u013272009/article/details/80807965)


## 例子主要文件介绍

文件名                    | 说明
------------------------ | --------------------
NOTES                    | 部署成功后，本文介绍如何使用redis
pv.yaml                  | 云存储配置。这里测试方便使用的hostpath。（注意：hostpath仅用于测试方便用）
Dockerfile               | 服务A 服务B 的Dockerfile
setup.sh                 | 安装脚本
values-production.yaml   | stable/redis 的配置参数


## 如何使用

1. 执行命令：

```shell
./setup.sh
```

2. 如何使用k8s redis

参考: [NOTES](NOTES)
