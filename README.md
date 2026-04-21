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

### Install gore (Go REPL)

Install `gore` and add `GOPATH` to `PATH` so go binaries can be executed from the terminal:
```
go install github.com/x-motemen/gore/cmd/gore@latest
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.zshrc
source ~/.zshrc
```
Now can enter `gore` REPL:
```
jordan@jordans-mbp go-tutorial % gore --version
gore 0.6.1 (rev: HEAD/go1.26.0)
jordan@jordans-mbp go-tutorial % which gore
/Users/jordan/go/bin/gore
jordan@jordans-mbp go-tutorial % gore
gore version 0.6.1  :help for help
gore> :help
    :import <package>     import a package
    :type <expr>          print the type of expression
    :print                print current source
    :write [<file>]       write out current source
    :clear                clear the codes
    :doc <expr or pkg>    show documentation
    :help                 show this help
    :quit                 quit the session
gore>
```
Each expression in the REPL is added to a Go program that is then executed whenever a new line is added.
In stdout you will see any Print statements as well as the return value of the last expression.

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

### `go run <package>`

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

### `go build <package>`

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

### `go build ./...`

If executed at the root of this project, multiple `main` packages will be matched. Go will compile them to verify they are error-free but discard the resulting binaries.

To compile one binary to be built and stored in the current directory, specify the directory of the main package, where there will be only 1 main package.

Alternatively, to store multiple binaries, specify a directory path to write all resulting executables:
```
jordan@jordans-mbp go-tutorial % go build -o ./bin ./...
go: cannot write multiple packages to non-directory ./bin
jordan@jordans-mbp go-tutorial % mkdir ./bin
jordan@jordans-mbp go-tutorial % go build -o ./bin ./...
jordan@jordans-mbp go-tutorial % ls ./bin
constant        hello           print
enum            logger          variable
jordan@jordans-mbp go-tutorial % rm -rf ./bin
jordan@jordans-mbp go-tutorial % ls ./bin
ls: ./bin: No such file or directory
jordan@jordans-mbp go-tutorial %
```

### `go install ./...`

To store multiple binaries in the go binary directory i.e. `$(go env GOPATH)/bin`:

```
jordan@jordans-mbp go-tutorial % go install ./...
jordan@jordans-mbp go-tutorial % ls $(go env GOPATH)/bin
constant        hello           print
enum            logger          variable
jordan@jordans-mbp go-tutorial % rm -rf $(go env GOPATH)/bin/*
jordan@jordans-mbp go-tutorial %
```

### `go install <package>@<version>`

Main packages (tools):
Installs package of particular version from an origin server into the go binary directory i.e. `$(go env GOPATH)/bin`. Additionally, the module and its dependencies are cached in the go module cache i.e. `$(go env GOPATH)/pkg`.

Library packages:
The module and its dependencies are cached in the go module cache. The code is compiled into an executable if it is imported into a main package that is compiled. If the code is imported into a library package, it is made part of the dependency tree via the go.mod file, but not compiled.

Tool example:
```
go install golang.org/x/tools/cmd/stringer@latest
```
This installs the `stringer` package into the go binary directory, as well as the entire `golang.org/x/tools` module into the go module cache:
```
jordan@Jordans-MacBook-Pro go-tutorial % ls -l $(go env GOPATH)
total 0
drwxr-xr-x  6 jordan  staff  192 Apr 16 11:27 bin
drwxr-xr-x  4 jordan  staff  128 Feb 20 12:14 pkg
jordan@Jordans-MacBook-Pro go-tutorial % ls -l $(go env GOPATH)/bin
total 155912
-rwxr-xr-x  1 jordan  staff  42279216 Feb 20 12:15 gopls
-rwxr-xr-x  1 jordan  staff  13044496 Apr 14 13:09 gore
-rwxr-xr-x  1 jordan  staff  16713664 Feb 20 12:15 staticcheck
-rwxr-xr-x@ 1 jordan  staff   7779936 Apr 16 11:27 stringer
jordan@Jordans-MacBook-Pro go-tutorial % ls -l $(go env GOPATH)/pkg
total 0
drwxr-xr-x  9 jordan  staff  288 Apr 14 13:09 mod
drwxr-xr-x  3 jordan  staff   96 Feb 20 12:14 sumdb
jordan@Jordans-MacBook-Pro go-tutorial % ls -l $(go env GOPATH)/pkg/mod 
total 0
drwxr-xr-x   4 jordan  staff  128 Feb 20 12:34 cache
drwxr-xr-x  16 jordan  staff  512 Apr 14 13:16 github.com
drwxr-xr-x   6 jordan  staff  192 Apr 14 13:09 go.lsp.dev
drwxr-xr-x   4 jordan  staff  128 Apr 14 13:09 go.uber.org
drwxr-xr-x   3 jordan  staff   96 Feb 20 12:14 golang.org
drwxr-xr-x   3 jordan  staff   96 Feb 20 12:14 honnef.co
drwxr-xr-x   4 jordan  staff  128 Feb 20 12:14 mvdan.cc
jordan@Jordans-MacBook-Pro go-tutorial % ls -l $(go env GOPATH)/pkg/mod/golang.org/x | grep tools
drwxr-xr-x   3 jordan  staff    96 Feb 20 12:14 tools
dr-xr-xr-x  26 jordan  staff   832 Apr 14 13:09 tools@v0.34.0
dr-xr-xr-x  25 jordan  staff   800 Feb 20 12:14 tools@v0.39.1-0.20260109155911-b69ac100ecb7
dr-xr-xr-x  25 jordan  staff   800 Feb 20 12:15 tools@v0.40.1-0.20260108161641-ca281cf95054
dr-xr-xr-x  25 jordan  staff   800 Feb 20 12:14 tools@v0.42.0
dr-xr-xr-x@ 25 jordan  staff   800 Apr 16 11:27 tools@v0.44.0
jordan@Jordans-MacBook-Pro go-tutorial % ls -l $(go env GOPATH)/pkg/mod/golang.org/x/tools@v0.44.0
total 56
-r--r--r--@  1 jordan  staff   913 Apr 16 11:27 CONTRIBUTING.md
-r--r--r--@  1 jordan  staff  1453 Apr 16 11:27 LICENSE
-r--r--r--@  1 jordan  staff  1303 Apr 16 11:27 PATENTS
-r--r--r--@  1 jordan  staff  3422 Apr 16 11:27 README.md
dr-xr-xr-x@  3 jordan  staff    96 Apr 16 11:27 benchmark
dr-xr-xr-x@  5 jordan  staff   160 Apr 16 11:27 blog
dr-xr-xr-x@ 28 jordan  staff   896 Apr 16 11:27 cmd
-r--r--r--@  1 jordan  staff    21 Apr 16 11:27 codereview.cfg
dr-xr-xr-x@  3 jordan  staff    96 Apr 16 11:27 container
dr-xr-xr-x@  4 jordan  staff   128 Apr 16 11:27 copyright
dr-xr-xr-x@  4 jordan  staff   128 Apr 16 11:27 cover
dr-xr-xr-x@ 14 jordan  staff   448 Apr 16 11:27 go
-r--r--r--@  1 jordan  staff   301 Apr 16 11:27 go.mod
-r--r--r--@  1 jordan  staff  1169 Apr 16 11:27 go.sum
dr-xr-xr-x@  3 jordan  staff    96 Apr 16 11:27 imports
dr-xr-xr-x@ 43 jordan  staff  1376 Apr 16 11:27 internal
dr-xr-xr-x@  4 jordan  staff   128 Apr 16 11:27 playground
dr-xr-xr-x@ 18 jordan  staff   576 Apr 16 11:27 present
dr-xr-xr-x@  7 jordan  staff   224 Apr 16 11:27 refactor
dr-xr-xr-x@  6 jordan  staff   192 Apr 16 11:27 txtar
jordan@Jordans-MacBook-Pro go-tutorial %
```
As demonstrated above, the module source code is copied into `$(go env GOPATH)/pkg/mod`. The modules are also referenced in `$(go env GOPATH)/pkg/mod/cache/download`, which is the raw network cache layer, where zip files are downloaded, module definitions, metadata, and available versions for tools are stored:
```
jordan@Jordans-MacBook-Pro go-tutorial % ls -l $(go env GOPATH)/pkg/mod/cache    
total 0
drwxr-xr-x  9 jordan  staff  288 Apr 14 13:09 download
-rw-r--r--@ 1 jordan  staff    0 Feb 20 12:34 lock
jordan@Jordans-MacBook-Pro go-tutorial % ls -l $(go env GOPATH)/pkg/mod/cache/download
total 0
drwxr-xr-x  17 jordan  staff  544 Apr 14 13:16 github.com
drwxr-xr-x   6 jordan  staff  192 Apr 14 13:09 go.lsp.dev
drwxr-xr-x   4 jordan  staff  128 Apr 14 13:09 go.uber.org
drwxr-xr-x   3 jordan  staff   96 Feb 20 12:14 golang.org
drwxr-xr-x   3 jordan  staff   96 Feb 20 12:14 honnef.co
drwxr-xr-x   4 jordan  staff  128 Feb 20 12:14 mvdan.cc
drwxr-xr-x   3 jordan  staff   96 Feb 20 12:14 sumdb
jordan@Jordans-MacBook-Pro go-tutorial % ls -l $(go env GOPATH)/pkg/mod/cache/download/golang.org/x/tools/@v | grep -E "v0.44.0|list"
-rw-r--r--@ 1 jordan  staff      135 Apr 16 11:27 list
-rw-r--r--@ 1 jordan  staff      192 Apr 16 11:27 v0.44.0.info
-rw-r--r--@ 1 jordan  staff        0 Apr 16 11:27 v0.44.0.lock
-rw-r--r--@ 1 jordan  staff      301 Apr 16 11:27 v0.44.0.mod
-rw-r--r--@ 1 jordan  staff  2711015 Apr 16 11:27 v0.44.0.zip
-rw-r--r--@ 1 jordan  staff       47 Apr 16 11:27 v0.44.0.ziphash
jordan@Jordans-MacBook-Pro go-tutorial %
```

Also existing is `$(go env GOPATH)/pkg/mod/cache/lock` and `$(go env GOPATH)/pkg/sumdb` and `$(go env GOPATH)/pkg/mod/cache/download/sumdb`.

The cache lock file is a concurrency lock for the module cache to prevent multiple go processes from corrupting cache writes, ensuring safe parallel downloads/builds during `go build`, `go test`, `gopls`, etc.

The `sumdb` is a database of checksums for modules, used to verify the integrity of downloaded modules. It is used by go to ensure that the modules downloaded are cryptographically consistent with the modules published, protecting against MITM attacks & code tampering.

The `latest` file is a cached checkpoint of the last query to sumdb. The lookup responses from the checksum database are cached in `$(go env GOPATH)/pkg/mod/cache/download/sumdb`. The real database is remote.
```
jordan@Jordans-MacBook-Pro go-tutorial % ls -l $(go env GOPATH)/pkg/sumdb
/sum.golang.org
total 8
-rw-r--r--@ 1 jordan  staff  188 Apr 16 11:27 latest
jordan@Jordans-MacBook-Pro go-tutorial % cat $(go env GOPATH)/pkg/mod/cache/download/sumdb/sum.golang.org/lookup/golang.org/x/tools@v0.44.0
52253027
golang.org/x/tools v0.44.0 h1:UP4ajHPIcuMjT1GqzDWRlalUEoY+uzoZKnhOjbIPD2c=
golang.org/x/tools v0.44.0/go.mod h1:KA0AfVErSdxRZIsOVipbv3rQhVXTnlU6UhKxHd1seDI=

go.sum database tree
52544510
+TGlZfnTizQCJgyZR56QixD8+X4dd8Wefy8U1w6k0Ig=

— sum.golang.org Az3grn6d1PSjcypDf9bfcFVU5fvJB7VkLc0MAbz8jqeQVnSbAj9pKcsB5dJt5VdswPDUtiFar5LPRBu+uEQAT1CDfAc=
jordan@Jordans-MacBook-Pro go-tutorial %
```

To summarize, these are the steps taken when a remote main package is installed e.g. `golang.org/x/tools/cmd/stringer@latest`:
1. Resolve version (module proxy / VCS)
2. Download → pkg/mod/cache/download
3. Verify → sumdb
4. Extract → pkg/mod/golang.org/x/tools@vX.Y.Z
5. Build ONLY cmd/stringer (package main)
6. Output binary → GOPATH/bin
The last 2 steps are skipped for library packages.

Note that a module proxy is a caching layer between the Go tool and the remote module registry e.g. GitHub. By default, Go may download modules from `proxy.golang.org`; if needed, it can also fetch them directly from the underlying version control system, based on the path `golang.org/x/tools` in this case.

Note that some modules have nested modules. In the below example, `gopls` is a module nested within the `tools` module:
```
jordan@Jordans-MacBook-Pro go-tutorial % ls -l $(go env GOPATH)/pkg/mod/golang.org/x | grep tools
drwxr-xr-x   3 jordan  staff    96 Feb 20 12:14 tools
dr-xr-xr-x  26 jordan  staff   832 Apr 14 13:09 tools@v0.34.0
dr-xr-xr-x  25 jordan  staff   800 Feb 20 12:14 tools@v0.39.1-0.20260109155911-b69ac100ecb7
dr-xr-xr-x  25 jordan  staff   800 Feb 20 12:15 tools@v0.40.1-0.20260108161641-ca281cf95054
dr-xr-xr-x  25 jordan  staff   800 Feb 20 12:14 tools@v0.42.0
dr-xr-xr-x@ 25 jordan  staff   800 Apr 16 11:27 tools@v0.44.0
jordan@Jordans-MacBook-Pro go-tutorial % ls -l $(go env GOPATH)/pkg/mod/golang.org/x/tools
total 0
dr-xr-xr-x  11 jordan  staff  352 Feb 20 12:14 gopls@v0.21.1
jordan@Jordans-MacBook-Pro go-tutorial %
```
The `gopls` code gets cached in the `tools` directory not the `tools@<version>` directories:
```
jordan@Jordans-MacBook-Pro go-tutorial % ls -l $(go env GOPATH)/pkg/mod/golang.org/x/tools/gopls@v0.21.1
total 64
-r--r--r--   1 jordan  staff   1453 Feb 20 12:14 LICENSE
-r--r--r--   1 jordan  staff    544 Feb 20 12:14 README.md
-r--r--r--   1 jordan  staff   1394 Feb 20 12:14 contributors.txt
dr-xr-xr-x  19 jordan  staff    608 Feb 20 12:14 doc
-r--r--r--   1 jordan  staff   1262 Feb 20 12:14 go.mod
-r--r--r--   1 jordan  staff  11538 Feb 20 12:14 go.sum
dr-xr-xr-x   3 jordan  staff     96 Feb 20 12:14 integration
dr-xr-xr-x  31 jordan  staff    992 Feb 20 12:14 internal
-r--r--r--   1 jordan  staff   1867 Feb 20 12:14 main.go
jordan@Jordans-MacBook-Pro go-tutorial % ls -l $(go env GOPATH)/pkg/mod/golang.org/x/tools@v0.44.0 | grep gopls
jordan@Jordans-MacBook-Pro go-tutorial %
```
See sourcecode here: https://github.com/golang/tools


### `go get <package>@<version>`

`go get` does not install packages, it just documents the projects dependency on them in `go.mod` and their versions & transitive dependency versions in `go.sum`.

You can globally install packages as described above, or you can run `go get` to associate pacakges with a project, in which case the packages and parent modules become cached locally, upon execution of `go build` / `go run` / `go test` / etc.

You would run `go get` when:
* add a library dependency after importing it globally
* upgrade a dependency to a newer version with `@<version>`
* pin a dependency to a specific version with `@<version>`
* remove a dependency with `@none`
* add a tool dependency with `-tool` (see next section below)
To add a library dependency before importing it globally, be implicit, see recommended approach below.

Example usage with https://github.com/gorilla/mux:
```
jordan@Jordans-MacBook-Pro go-tutorial % go get github.com/gorilla/mux@latest
go: downloading github.com/gorilla/mux v1.8.1
go: added github.com/gorilla/mux v1.8.1
jordan@Jordans-MacBook-Pro go-tutorial % go mod tidy
jordan@Jordans-MacBook-Pro go-tutorial %
```
You can pass `-u` to `go get` to upgrade ALL dependencies to their latest versions, which is generally not recommended, as it could cause unsuspecting breaking changes.

The recommended approach is actually to import a module like `import "github.com/gorilla/mux"` and then run `go mod tidy`. This will automatically download the module and its transitive dependencies, without explicitly executing `go get`. See ./masterclass/section4/dependency for mux import and then:
```
jordan@Jordans-MacBook-Pro go-tutorial % go mod tidy
go: finding module for package github.com/gorilla/mux
go: found github.com/gorilla/mux in github.com/gorilla/mux v1.8.1
jordan@Jordans-MacBook-Pro go-tutorial %
```

### `go get -tool <tool>`

Library vs Tool comparison:
| Aspect             | Library                | Tool                |
| ------------------ | ---------------------- | ------------------- |
| Package type       | normal (`package mux`) | `package main`      |
| Used via           | `import`               | CLI / `go generate` |
| Included in binary | ✅ yes                  | ❌ no                |
| Stored in `go.mod` | ✅ yes                  | optional (`tool`)   |
| Installed with     | `go mod tidy`          | `go install`        |
| Runs when          | build/run              | manually / generate |


To add a tool dependency such as `stringer`:
```
go get -tool golang.org/x/tools/cmd/stringer
go mod tidy
```
Example:
```
jordan@Jordans-MacBook-Pro go-tutorial % go get -tool golang.org/x/tools/cmd/stringer
go: added golang.org/x/mod v0.35.0
go: added golang.org/x/sync v0.20.0
go: added golang.org/x/tools v0.44.0
jordan@Jordans-MacBook-Pro go-tutorial % go mod tidy
go: downloading github.com/google/go-cmp v0.6.0
jordan@Jordans-MacBook-Pro go-tutorial %
```
This updates `go.mod` and `go.sum` (which document the tool dependency and its transitive dependencies):
```
module github.com/chord-memory/go-tutorial

go 1.25.0

tool golang.org/x/tools/cmd/stringer

require (
	golang.org/x/mod v0.35.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/tools v0.44.0 // indirect
)
```
```
github.com/google/go-cmp v0.6.0 h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=
github.com/google/go-cmp v0.6.0/go.mod h1:17dUlkBOakJ0+DkrSSNjCkIjxS6bF9zb3elmeNGIjoY=
golang.org/x/mod v0.35.0 h1:Ww1D637e6Pg+Zb2KrWfHQUnH2dQRLBQyAtpr/haaJeM=
golang.org/x/mod v0.35.0/go.mod h1:+GwiRhIInF8wPm+4AoT6L0FA1QWAad3OMdTRx4tFYlU=
golang.org/x/sync v0.20.0 h1:e0PTpb7pjO8GAtTs2dQ6jYa5BWYlMuX047Dco/pItO4=
golang.org/x/sync v0.20.0/go.mod h1:9xrNwdLfx4jkKbNva9FpL6vEN7evnE43NNNJQ2LF3+0=
golang.org/x/tools v0.44.0 h1:UP4ajHPIcuMjT1GqzDWRlalUEoY+uzoZKnhOjbIPD2c=
golang.org/x/tools v0.44.0/go.mod h1:KA0AfVErSdxRZIsOVipbv3rQhVXTnlU6UhKxHd1seDI=

```
See details on `go.mod` below.

You can see installed tools like this:
```
jordan@Jordans-MacBook-Pro go-tutorial % go tool    
asm
cgo
compile
cover
fix
link
preprofile
vet
golang.org/x/tools/cmd/stringer
jordan@Jordans-MacBook-Pro go-tutorial %
```

### `go mod init`

`go mod init <repo>` to generate go.mod file which tracks your code dependencies. If you publish a module, use the path to your code's repository. 
```
jordan@Jordans-MBP go-tutorial % go mod init github.com/chord-memory/go-tutorial 
go: creating new go.mod: module github.com/chord-memory/go-tutorial
go: to add module requirements and sums:
        go mod tidy
jordan@Jordans-MBP go-tutorial %
```
At first the go.mod file may be tiny:
```
module github.com/chord-memory/go-tutorial

go 1.25.0

```
Then over time it may grow to include:
* `require` directives for modules your code imports
* `tool` directives for dev tools like `stringer`
* occasionally `replace` for local development
* rarely `exclude` if you must avoid a bad version

In normal day-to-day use, you usually do not hand-edit it much. The common workflow is:
* import a package, then run `go build` / `go test` / `go mod tidy`
* or run `go get ...` to add or upgrade a dependency
* or run `go get -tool ...` to add a tool dependency
* then commit both `go.mod` and `go.sum`

Explanations of the `git mod init` command:
* [Reddit: Can someone please dumb down go mod init for me?](https://stackoverflow.com/questions/67606062/can-someone-please-dumb-down-go-mod-init-for-me)
* [go.dev: Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module)
* [go.dev: Managing Dependencies](https://go.dev/doc/modules/managing-dependencies#naming_module)

### `go generate`

# TODO

### `go fmt`

`go fmt` to format according to convention.
```
jordan@Jordans-MBP go-tutorial % go fmt masterclass/section1/hello.go
masterclass/section1/hello.go
jordan@Jordans-MBP go-tutorial %
```
`gofmt` is the underlying tool that `go fmt` uses.
`gofmt` is file-level while `go fmt` is package-level.
`gofmt` is used by scripting editors while `go fmt` is used by CI or for repo cleanup.
`goimports` is more common bc it does `gofmt` and fixes impots.
`gopls` is used by IDE for formatting, autocompletion, references, etc.

### `go test`

`go test`
```
```

### `golangci-lint`

`golangci-lint` is a third party linter that must be installed separately from the main go binary.

## PostgreSQL

Use pgAdmin 4 as GUI to interact with db
```
brew install --cask pgadmin4
```