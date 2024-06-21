# 第一阶段：构建二进制文件
FROM golang:1.22.2 AS builder

ENV GOPROXY https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码和其他所有文件
COPY . .

# 构建每个微服务的二进制文件
ARG SERVICE
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/${SERVICE} ./app/${SERVICE}/cmd/main.go

## 第二阶段：创建运行时镜像
#FROM scratch
#
## 设置构建参数
#ARG SERVICE
#
## 复制所有文件，包括构建好的二进制文件和配置文件
#COPY --from=builder /app /app
#COPY --from=builder /out/${SERVICE} /${SERVICE}
#
## 设置工作目录
#WORKDIR /app

# 设置入口点和命令
ARG SERVICE
CMD ["/gateway"]

