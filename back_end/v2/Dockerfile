FROM golang:1.22.2 AS builder

ENV GOPROXY https://goproxy.cn,direct
# 设置工作目录
WORKDIR /app

# 复制项目文件
COPY . .

# 编译应用程序
RUN go build -o database_design .

FROM scratch

# 从第一阶段复制编译好的二进制文件和相关文件到当前镜像中
COPY --from=builder /app/database_design /app/database_design
COPY settings.yaml /app/settings.yaml

# 设置工作目录
WORKDIR /app

EXPOSE 4000
ENTRYPOINT ["/database_design"]
