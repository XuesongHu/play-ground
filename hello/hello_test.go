package hello

import (
	"context"
	"testing"
	"time"

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

func TestServer(t *testing.T) {
	// set up the server, client and server use name to identify
	// each other, so both client and server need to create a
	// micro service object with the same name
	service := micro.NewService(
		micro.Name("greeter"),
	)
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	go service.Run()
	time.Sleep(5 * time.Second)
	client := micro.NewService(
		micro.Name("greeter"),
	)
	greeter := proto.NewGreeterService("greeter", client.Client())
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		t.Errorf("got %q and want nil", err)
	}
	want := "Hello John"
	got := rsp.GetGreeting()
	if got != want {
		t.Errorf("got %q and want %q", got, want)
	}
}
