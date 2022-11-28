package main

import (
	"go-micro-service-template/service"
	"log"
	"os"
)

func main() {
	addrs := ":3000"
	svc := service.NewRandomJokeService("https://api.api-ninjas.com/v1/dadjokes?limit=1", os.Getenv("API_NINJAS"))
	svc = service.NewLoggingService(svc)

	apiServer := service.NewApiServer(svc)
	log.Fatal(apiServer.Start(addrs))
}
