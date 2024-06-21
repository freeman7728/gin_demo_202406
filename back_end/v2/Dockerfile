# 第一阶段：构建阶段，使用 golang:1.22.2 作为基础镜像，命名为 builder
FROM golang:1.22.2 AS builder

# 设置环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

# 设置工作目录
WORKDIR /app

# 复制并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod tidy

# 将整个项目的代码复制到容器中
COPY . .

# 构建应用程序
RUN go build -o database_design .

# 第二阶段：生成最终的镜像，基于 scratch 空镜像构建
FROM scratch

# 从第一阶段复制构建好的二进制文件和相关文件到当前镜像中
COPY --from=builder /app/database_design /database_design

# 设置工作目录
WORKDIR /

# 声明端口
EXPOSE 4000

# 设置容器启动命令
ENTRYPOINT ["/database_design"]
