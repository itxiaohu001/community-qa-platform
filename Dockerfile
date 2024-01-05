# 使用一个基础镜像，可以选择适合你项目的基础镜像
FROM ubuntu:latest

# 设置工作目录
WORKDIR /app

# 复制可执行文件和config.json到容器中
COPY my-app /app/
COPY config.json /app/

# 启动应用程序
CMD ["./my-app", "-c", "config.json"]
