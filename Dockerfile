# 编译阶段：引用最小编译环境
FROM golang:1.22.0-alpine AS builder
ENV TZ=Asia/Shanghai

# 为了减小镜像大小，添加必要的编译工具
RUN apk add --no-cache tzdata build-base

# 镜像默认工作目录
WORKDIR /build

# 明确指定需要的文件，而不是复制所有代码，提高构建效率
COPY go.mod go.sum ./

# 配置镜像golang的默认配置,方便拉取依赖
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

# 构建
COPY . .
RUN GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o server .

# 构建阶段：使用 alpine 最小构建
FROM alpine:3.19
ENV TZ=Asia/Shanghai

# 设置镜像工作目录
WORKDIR /app

# 在builder阶段复制可执行的go二进制文件和配置文件
COPY --from=builder /build/server .
COPY --from=builder /build/config/ config

# 确保运行时所需的库已安装
RUN apk --no-cache add ca-certificates

# 启动服务器
CMD ["./server"]

# 开放端口
EXPOSE 8080