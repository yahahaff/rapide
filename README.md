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
### Run
```shell
cd path/to/your/project
go mod tidy
swag init
air
```