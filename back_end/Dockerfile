# 第一阶段：构建二进制文件
FROM golang:1.22.2 AS builder

ENV GOPROXY https://goproxy.cn,direct

# 设置工作目录
WORKDIR /

# 复制 go.mod 和 go.sum 文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建每个微服务的二进制文件
ARG SERVICE
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/${SERVICE} ./app/${SERVICE}/cmd/main.go

# 第二阶段：创建运行时镜像
FROM scratch

# 将构建好的二进制文件从构建阶段复制到运行时镜像中
ARG SERVICE
COPY --from=builder /out/${SERVICE} /${SERVICE}

# 设置入口点和命令
CMD ["./employee"]
