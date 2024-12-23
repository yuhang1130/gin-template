# 基础镜像，基于golang的alpine镜像构建--编译阶段
FROM golang:1.23.4-alpine AS builder

# 设置工作目录
WORKDIR /build

# 明确指定需要的文件，而不是复制所有代码，提高构建效率
COPY go.mod go.sum ./

# 配置镜像golang的默认配置,方便拉取依赖
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

# 构建
COPY . .
RUN GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build main.go
# RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main.go


# 使用alpine这个轻量级镜像为基础镜像--运行阶段
FROM alpine:latest AS runner

# 设置镜像工作目录
WORKDIR /app

# 设置环境变量
ENV env=test

# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /build/main .
COPY --from=builder /build/config/ ./config
# 将时区设置为东八区
RUN echo "https://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories \
  && echo "https://mirrors.aliyun.com/alpine/v3.8/community/" >> /etc/apk/repositories \
  && apk add --no-cache tzdata \
  && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
  && echo Asia/Shanghai > /etc/timezone \
  && apk del tzdata

# 需暴露的端口
EXPOSE 8080

# 启动服务器
# docker run命令触发的真实命令(相当于直接运行编译后的可运行文件)
ENTRYPOINT ["./main"]