FROM golang:1.22.2 AS builder
# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64\
    GOPROXY=https://goproxy.cn
# 设置/usr/local/goproject，进入容器就会直接进入到这个目录下,而不是进入到默认根目录下面
WORKDIR .
# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod tidy
# 将代码复制到容器中
COPY . .
RUN go build -o database_design .

FROM scratch

# 复制所有文件，包括构建好的二进制文件和配置文件
COPY --from=builder . .
# 设置工作目录
WORKDIR .
EXPOSE 4000
# 设置入口点和命令
ENTRYPOINT ["/database_design"]

