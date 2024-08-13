# Go Project Template

## Overview

**Go Project Template** is a streamlined scaffold for rapidly bootstrapping Go projects. It's designed to accelerate the development process by providing a solid foundation for implementing business logic efficiently.

## Key Features

- **Dependency Injection** with `wire`.
- **HTTP Server** setup using the `gin` framework.
- **Configuration Management** via `viper`.
- **Database Support** for both `sqlite3` and `mysql`, with SQL queries handled by `didi/gendry`.
- **Logging** using the `slog` standard library.
- **Docker Integration** with a `Dockerfile` for easy containerization, including optimized image sources for users in China.

## Getting Started

The core business logic is organized in the following directories:

- `data` for database queries.
- `biz` for business logic.
- `service` for request handling and parameter processing.
- `service/binding.go` for route registration.

## Quick Start

Clone the repository and start the server with Docker:

```shell
git clone https://github.com/zoujiejun/go-project-template.git
cd go-project-template
docker build -t go-project-template .
docker run --rm -p 8080:8080 go-project-template
```

Test the endpoints with the following `curl` commands:

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
## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
