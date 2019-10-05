package hello

import (
	"context"

	proto "github.com/XuesongHu/play-ground/hello/proto"
	"github.com/micro/go-micro"
)

// Hello return string
func Hello() string {
	service := micro.NewService()
	return service.String()
}

// Greeter type to be used by server
type Greeter struct{}

// Hello implements interface for server side
func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}
