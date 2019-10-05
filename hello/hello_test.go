package hello

import (
	"context"
	"testing"

	proto "github.com/XuesongHu/play-ground/hello/proto"
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
	if got != nil {
		t.Errorf("got %q and want nil", got)
	}
	res := rsp.GetGreeting()
	want := "Hello John"
	if res != want {
		t.Errorf("got %q and want %q", res, want)
	}
}
