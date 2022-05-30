# gRPC Server -> Client; Calculator Example

This repo contains an example [`gRPC`](https://grpc.io/docs/what-is-grpc/introduction/) server and client implemented in a form of `calculator` service that would perform basic math methods:

- Addition
- Subtraction
- Division
- Multiplication

The results are printed to `STDOUT`

> NOTE: This is un-opinionated example, therefore useful utils, features and third-packages (buf, viper, server deployment, clang, grcp-health-probe, envs) were not used/implemented for the sake of simplicity.

## Prerequisites

- Go **1.18**
- Protocol Buffers Compiler, `protoc` (`3.19.4` to date version):
  - Homebrew: `brew install protobuf`
  - Alpine 3.15: `apk add protobuf-dev protobuf`
  - Ubuntu 21.10: `apt-get install protobuf-compiler libprotobuf-dev`
- Protocol Buffer Plugin for Go:
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0`
- gRPC Plugin for Go:
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0`

## Usage

> Please make sure that prerequisites are met before usage

![ezgif com-gif-maker (1)](https://user-images.githubusercontent.com/81821185/171068756-99cb4e4e-03c2-48c5-9b97-718f3d376f46.gif)

From the module root:

1. Build the binary files: `make build`

```bash
❯ make build
==> Building binaries...
==> Done ./bin/
```

2. Start server: `./bin/server`

```bash
❯ ./bin/server
2022/05/30 17:20:14 server listening at 127.0.0.1:8080
```

3. Call client with args: `./bin/client -method add -a 17 -b 24`
   > `add`, `sub`, `mul`, `div`

```bash
❯ ./bin/client -method add -a 17 -b 24
2022/05/30 20:35:34 Result: 41
```

File `tree` structure

```
├── Makefile // automated utilities
├── README.md // you are here
├── bin // binaries
│   ├── client
│   └── server
├── cmd // main applications for this project
│   ├── client
│   │   ├── client.go
│   │   └── client_test.go
│   └── server
│       ├── server.go
│       └── server_test.go
├── coverage.txt // generated test coverage
├── gen // generated go .proto files
│   ├── calc.pb.go
│   └── calc_grpc.pb.go
├── go.mod
├── go.sum
└── pkg
    └── calc // calc project .proto file
        └── calc.proto
```

## Utils

### To list all Makefile helper commands use: `make help`

```bash
❯ make generate
==> Generated proto files in ./gen
```

### To generate new Go Proto files use: `make generate`

```bash
❯ make help
build                          Build the binary files
clean                          Remove previous build
dep                            Get the dependencies
...
...

```

### To invoke test cases use: `make test`

```bash
❯ make test
=== RUN   TestCheckMethodExist
--- PASS: TestCheckMethodExist (0.00s)
=== RUN   TestCheckMethodExistCrashes
--- PASS: TestCheckMethodExistCrashes (0.01s)
PASS
...
...

```
