To generate go code from protobuf

`cd hello/proto`

`go get github.com/golang/protobuf/{proto,protoc-gen-go}`

`go get github.com/micro/protoc-gen-micro`

`protoc --proto_path=. --micro_out=. --go_out=. hello.proto`