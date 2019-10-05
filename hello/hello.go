package hello

import "github.com/micro/go-micro"

// Hello return string
func Hello() string {
	service := micro.NewService()
	return service.String()
}
