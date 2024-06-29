# 超市进销存管理系统

本项目使用Gin框架，实现了若干增删改查接口

使用Redis与SMTP实现邮箱验证服务

v1版本为Go-micro+Etcd+RPC实现微服务，使用Gin作为网关，最终因为服务器资源消耗过大未能完成线上部署

v2版本为单体架构，使用Gin框架完成

项目主要技术栈如下

* Gin
* Gorm
* Go-micro(v1)
* Swagger
* RPC(v1)
* Mysql
* Redis
* Etcd(v1)
* Docker
* Docker-compose
* Nginx

## 后端部署

在v2下执行命令

```shell
docker-compose up -d --build
```

服务运行在

```shell
localhost:4000
```

可以在settings.yaml中修改

```yaml
system:
  host: "0.0.0.0"
  port: "4000"
```

Mysql容器在建立时有若干初始数据，但未进行数据持久化，需自行完成容器卷挂载

## 前端部署

编辑nginx.conf文件

替换server内部的location，如下

```shell
location / {
	    root  /root/database_design/database-course-design/front_end/dist;
	    index index.html;
	    try_files $uri $uri/ /index.html;
	}
```

