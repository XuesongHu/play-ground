package main

import (
	"context"
	"fmt"
	"os"

	"github.com/XuesongHu/play-ground/hello"
	proto "github.com/XuesongHu/play-ground/hello/proto"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func runClient(service micro.Service) {
	greeter := proto.NewGreeterService("greeter", service.Client())
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.GetGreeting())
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Flags(cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
	)
	service.Init(micro.Action(func(c *cli.Context) {
		if c.Bool("run_client") {
			runClient(service)
			os.Exit(0)
		}
	}))

	proto.RegisterGreeterHandler(service.Server(), new(hello.Greeter))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
