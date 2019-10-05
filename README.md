To generate go code from protobuf

`cd hello/proto`

The following two steps are optional, only needed if you don't have those packages installed.

`go get github.com/golang/protobuf/{proto,protoc-gen-go}`

`go get github.com/micro/protoc-gen-micro`

`protoc --proto_path=. --micro_out=. --go_out=. hello.proto`

The folder structure for the project is as follows:
```
.
├── README.md
├── go.mod
├── go.sum
├── hello
│   ├── hello.go
│   ├── hello_test.go
│   └── proto
│       ├── hello.micro.go
│       ├── hello.pb.go
│       └── hello.proto
└── main.go
```