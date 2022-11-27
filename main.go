package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	apiKey := os.Getenv("API_NINJAS")
	fmt.Println(apiKey)
	// NewRandomJokeService("https://api.api-ninjas.com/v1/dadjokes?limit=1", apiKey)
	// svc = NewLoggingService(svc)

	// joke, err := svc.GetRandomJoke(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v", joke)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.api-ninjas.com/v1/dadjokes?limit=1", nil)
	req.Header.Set("X-Api-Key", apiKey)
	res, _ := client.Do(req)

	defer res.Body.Close()

	fmt.Println(res.StatusCode)
	if res.StatusCode == http.StatusOK {
		// jokes := &service.DecodedRandomJoke{}

		// json.NewDecoder(res.Body).Decode(jokes)
		// fmt.Printf("%+v", jokes)

		bodyBytes, _ := io.ReadAll(res.Body)

		bodyString := string(bodyBytes)
		fmt.Printf("%+v", bodyString)
	}

}
