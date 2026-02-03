FROM golang:1.19

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0

# 构建 Rapide 应用程序
RUN go mod tidy
RUN go build -ldflags="-s -w" -o rapide .

# 暴露端口，如果 Rapide 应用程序监听了特定的端口，请确保与其一致
EXPOSE 8080

# 设置容器启动命令，默认启动 Rapide 应用程序
CMD ["./rapide"]
