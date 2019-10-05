package main

import (
	"fmt"

	"github.com/XuesongHu/play-ground/hello"
)

func main() {
	got := hello.Hello()
	fmt.Println(got)
	fmt.Println("Hello, world! from main")
}
