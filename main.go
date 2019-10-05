package main

import (
	"fmt"

	"ibm.com/hello/hello"
)

func main() {
	got := hello.Hello()
	fmt.Println(got)
	fmt.Println("Hello, world! from main")
}
