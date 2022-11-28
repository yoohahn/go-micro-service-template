package main

import (
	"context"
	"fmt"
	"go-micro-service-template/service"
	"log"
	"os"
)

func main() {
	svc := service.NewRandomJokeService("https://api.api-ninjas.com/v1/dadjokes?limit=1", os.Getenv("API_NINJAS"))
	svc = service.NewLoggingService(svc)
	joke, err := svc.GetRandomJoke(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n<--->")
	fmt.Println("")
	fmt.Printf("%v", joke)
}
