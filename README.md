<div align="center">
<br/>
<h1>Rapide</h1>
</div>

<div align="center">
  <img src="https://github.com/yahahaff/rapide/workflows/Go/badge.svg" alt="CI/CD Badge">
  <img src="https://img.shields.io/github/license/yahahaff/rapide?style=flat-square" alt="GitHub License">
  <img src="https://img.shields.io/github/go-mod/go-version/yahahaff/rapide" alt="GitHub go.mod Go version">
  <a href="https://sourcegraph.com/github.com/yahahaff/rapide/-/tree/codec?badge"><img src="https://sourcegraph.com/github.com/yahahaff/rapide/-/badge.svg?v=4" alt="Sourcegraph"></a>
  <a href="https://codecov.io/gh/yahahaff/rapide"><img src="https://codecov.io/gh/yahahaff/rapide/branch/main/graph/badge.svg?v=4" alt="codecov"></a>
  <a href="https://pkg.go.dev/github.com/yahahaff/rapide/codec"><img src="https://pkg.go.dev/badge/github.com/yahahaff/rapide/codec.svg" alt="Go Reference"></a>
  <a href="https://goreportcard.com/report/github.com/yahahaff/rapide/codec"><img src="https://goreportcard.com/badge/github.com/yahahaff/rapide?v=4" alt="rcard"></a>

</div>
👋 Welcome to Rapide!

`Rapide` is a simple backend project framework developed in Go, with features covering the most popular technology stacks. It's highly suitable for beginners looking to get started with learning.

---
### Install air
`Air` is a hot-reloading tool for Go. It can monitor changes to files or directories, automatically compile, and restart the program, thereby improving development efficiency. It requires Go version 1.16 or higher.
 ```shell
 go install github.com/cosmtrek/air@latest
 ```
### Run on IDE
```shell
git clone https://github.com/yahahaff/rapide.git
cd rapide
go mod tidy
air
```
### Envs
| 变量名                        | 默认值         | 简介                      |
|----------------------------|-------------|-------------------------|
| **APP_ENV**                | debug       | debug,test,release      |
| **APP_PORT**               | 8000        | app port                |
| **DB_DRIVER**              | sqlite      | 数据库连接驱动器 支持mysql,sqlite |
| **DB_CONNECTION_HOST**     | localhost   | mysql主机地址               |
| **DB_CONNECTION_PORT**     | 3306        | mysql数据库端口              |
| **DB_CONNECTION_USERNAME** | root        | mysql数据库用户              |
| **DB_CONNECTION_PASSWORD** | password    | mysql数据库密码              |
| **DB_CONNECTION_DATABASE** | rapide      | mysql数据库                |
| **DB_CONNECTION_FILE**     | database.db | sqlite db file          |
| **REDIS_HOST**             | 8000        | redis host              |
| **REDIS_PORT**             | 6379        | redis port              |
| **LOG_PATH**               | rapide.log  | 日志路径                    |
