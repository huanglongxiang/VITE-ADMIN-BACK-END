# VITE-ADMIN-BACK-END
vite-admin后端系统
- 基于go语言开发，使用go-zero框架、gorm作为数据库连接、gorm-gen作为代码生成、redis、swagger作为接口文档生成
``
- deploy 项目部署相关信息，如部署时一些程序的配置，sql后者dockerfile
- pkg 项目的公共目录****
- docker-compose.yml docker compose 文件
- Makefile 项目编译脚本程序
  #先将程序编译成C进制可执行的文
  GOOS=linux GoARCH=amd64 CGO_ENABLED=0 go build -o bin/user-rpc ./apps/user/rpc/user.go
  #然后根据二进制文件构建成镜像文件
  docker build -t user-rpc -f ./dockerfile_rpc .
  再修改构建的镜像标签#
  $ docker tag user-rpc registry.cn-hangzhou.aliyuncs.com/easy-im/user-rpc-test:latest
  #然后推送到阿里云上
  docker push registry.cn-hangzhou.aliyuncs.com/easy-im/user-rpc-test:1atest
  #在部署的时候拉取下来构建容器运行即可
``