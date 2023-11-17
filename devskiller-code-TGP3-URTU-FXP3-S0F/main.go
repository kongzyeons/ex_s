package main

import (
	"fmt"
	"log"

	"gitlab.com/devskiller-tasks/rest-api-blog-golang/bootstrap"
)

func main() {
	defaultPort := 8080
	fmt.Println(defaultPort)
	if err := bootstrap.Init(defaultPort); err != nil {
		log.Fatalf("Service will be shutdown because error ocurred:  %+v", err.Error())
	}
}
