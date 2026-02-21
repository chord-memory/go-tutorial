# Go Tutorial

This a workspace to organize my notes and code related to the following Go courses:
* [Go (Golang) Masterclass: Learn Like a Google Engineer by Joseph Abah](https://www.udemy.com/course/learn-golang-like-google-engineers-do/)
* [Complete Microservices with Go by Tiago Taquelim](https://www.udemy.com/course/complete-microservices-with-go/)

## Getting Started

### Install Go

```
brew install go
go version
```

### Definitions

module = see module name in `go.mod`, `go.mod` location defines module boundary, usually 1 module per repo, initialized with `go init mod <name>`, module name + package name used to import packages

package = a directory that can be imported like `import "github.com/chord-memory/go-tutorial/masterclass/section2/mathutils"`, the go files within the package specify `package mathutils`

executable = a go file with `package main` and `func main`, when `go build` executes, these compiled files are generated, only 1 per package (can have multiple `package main` but only 1 `func main`) (code in same package does not need to be imported)

Every directory is either:
* A library package (importable)
* Or a main package (executable)

Note: Subdirectories are different packages.

### Example Directory Layout

```
cmd/
  api/
    main.go
    server.go
    routes.go
  worker/
    main.go
    jobs.go
    queue.go
internal/
  db/
    db.go
  auth/
    auth.go
```
Each `cmd/*` directory is a separate executable.

Each executable has:
* 1 main() function
* Many supporting files (optionally)
* Imports shared libraries (`internal/*`)

## Go CLI

### `go run`

`go run` compiles a program into temporary directory and executes the program. (program = package)
```
jordan@Jordans-MBP go-tutorial % go run masterclass/section1/hello/main.go
Hello, world!
jordan@Jordans-MBP go-tutorial % go run ./masterclass/section1/hello      
Hello, world!
jordan@Jordans-MBP go-tutorial % cd masterclass/section1/hello 
jordan@Jordans-MBP hello % go run . 
Hello, world!
jordan@Jordans-MBP hello %
```

Files don’t require `./`. Directories do (to prove relative local path and not standard library package or a module in your module cache).

However, people almost always run directories because real programs are multi-file. Running a single file ignores other files in the package.

### `go build`

`go build` compiles program into current directory. `hello` in the example below is a compiled program.
```
jordan@Jordans-MBP go-tutorial % ls
README.md       go.mod          masterclass
jordan@Jordans-MBP go-tutorial % go build ./masterclass/section1/hello
jordan@Jordans-MBP go-tutorial % ls
README.md       go.mod          hello           masterclass
jordan@Jordans-MBP go-tutorial % ./hello
Hello, world!
jordan@Jordans-MBP go-tutorial % cd masterclass/section1/hello 
jordan@Jordans-MBP hello % go build .
jordan@Jordans-MBP hello % ls
hello   main.go
jordan@Jordans-MBP hello %
```

### `go run .` & `go build .`

// TODO when there are more than 1 executable

### `go mod init`

`go mod init <repo>` to generate go.mod file which tracks your code dependencies. If you publish a module, use the path to your code's repository. 
```
jordan@Jordans-MBP go-tutorial % go mod init github.com/chord-memory/go-tutorial 
go: creating new go.mod: module github.com/chord-memory/go-tutorial
go: to add module requirements and sums:
        go mod tidy
jordan@Jordans-MBP go-tutorial %
```
Explanations of the `git mod init` command:
* [Reddit: Can someone please dumb down go mod init for me?](https://stackoverflow.com/questions/67606062/can-someone-please-dumb-down-go-mod-init-for-me)
* [go.dev: Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module)
* [go.dev: Managing Dependencies](https://go.dev/doc/modules/managing-dependencies#naming_module)

### `go fmt`

`go fmt` to format according to convention.
```
jordan@Jordans-MBP go-tutorial % go fmt masterclass/section1/hello.go
masterclass/section1/hello.go
jordan@Jordans-MBP go-tutorial %
```

### `go test`

`go test`
```
```

## `go install` & `go get`

`go install` and `go get`
```
```


### `golangci-lint`

`golangci-lint` is a third party linter that must be installed separately from the main go binary.

## PostgreSQL

Use pgAdmin 4 as GUI to interact with db
```
brew install --cask pgadmin4
```