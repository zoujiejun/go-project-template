# go-project-template

## 概览

**go-project-template** 是一个简洁的项目脚手架，旨在帮助开发者快速启动Go项目，并高效实现业务逻辑。

## 核心特性

- **依赖注入**：使用 `wire` 实现。
- **HTTP 服务器**：基于 `gin` 框架构建。
- **配置文件加载**：通过 `viper` 管理。
- **数据库支持**：支持 `sqlite3` 和 `mysql`，SQL 查询基于 `didi/gendry` 实现。
- **日志记录**：使用 `slog` 标准库。
- **Docker 支持**：包含 `Dockerfile`，支持容器化部署，特别为中国用户提供了加速镜像。

## 使用方法

业务逻辑主要存放在以下目录：

- `data`：数据库查询。
- `biz`：业务逻辑编写。
- `service`：接收请求和参数处理。
- `service/binding.go`：路由注册。

## 快速开始

克隆仓库并使用 Docker 启动服务：

```shell
git clone https://github.com/zoujiejun/go-project-template.git
cd go-project-template
docker build -t go-project-template .
docker run --rm -p 8080:8080 go-project-template
```

使用以下 `curl` 命令测试端点：

```shell
curl -X POST 'localhost:8080/foo' \
-H 'Content-Type: application/json' \
-d '{"name": "hello"}'
```

```shell
curl -X GET 'localhost:8080/foo/sample/list'
```

```shell
curl -X GET 'localhost:8080/foo/1'
```

## 开源许可证

本项目采用 MIT 许可证。

