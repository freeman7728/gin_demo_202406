# 第一阶构建
FROM golang:alpine

ENV GOPROXY https://goproxy.cn,direct

RUN mkdir /app/

ADD . /app/

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/main.go -o main .

# 第二阶构建，从空白镜像开始
FROM scratch

## 复制上层构建的文件
COPY --from=0 /app/main /

CMD ["./main"]
