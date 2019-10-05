package hello

import (
	"context"
	"testing"

	proto "github.com/XuesongHu/play-ground/hello/proto"
	"github.com/micro/go-micro"
)

// TestHello tests Hello()
func TestHello(t *testing.T) {
	got := Hello()
	want := "micro"
	if got != want {
		t.Errorf("got %q and want %q", got, want)
	}
}

func TestHelloServer(t *testing.T) {
	g := Greeter{}
	ctx := context.Background()
	req := new(proto.HelloRequest)
	req.Name = "John"
	rsp := new(proto.HelloResponse)
	got := g.Hello(ctx, req, rsp)

	// assert return value is nil
	if got != nil {
		t.Errorf("got %q and want nil", got)
	}
	res := rsp.GetGreeting()
	want := "Hello John"
	// assert rsp object is updated
	if res != want {
		t.Errorf("got %q and want %q", res, want)
	}
}

// TestHelloClient tests client without connection
func TestHelloClient(t *testing.T) {
	service := micro.NewService()
	greeter := proto.NewGreeterService("greeter", service.Client())
	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	// there should be an error since there is no connection
	if err == nil {
		t.Errorf("got %q and want not nil", err)
	}
	want := ""
	got := rsp.GetGreeting()
	// the actual should be empty as the call could not go through
	if got != want {
		t.Errorf("got %q and want %q", got, want)
	}
}
