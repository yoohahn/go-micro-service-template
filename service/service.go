package service

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetRandomJoke(context.Context) (*ApiBody, error)
}

type RandomJokeService struct {
	url    string
	apiKey string
}

func NewRandomJokeService(url string, key string) Service {
	return &RandomJokeService{
		url:    url,
		apiKey: key,
	}
}

func (rjs *RandomJokeService) GetRandomJoke(ctx context.Context) (*ApiBody, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.api-ninjas.com/v1/dadjokes?limit=1", nil)
	req.Header.Set("X-Api-Key", rjs.apiKey)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	/*
		// Debug printing the response
		bodyBytes, _ := io.ReadAll(res.Body)
		bodyString := string(bodyBytes)
		fmt.Printf("%+v", bodyString)
	*/

	apiBody := &ApiBody{
		Code:      res.StatusCode,
		HasErrors: res.StatusCode != http.StatusOK,
	}

	var jokes []RandomJoke
	if err := json.NewDecoder(res.Body).Decode(&jokes); err != nil {
		apiBody.Errors = append(apiBody.Errors, err.Error())
		return apiBody, err
	}

	if len(jokes) > 0 {
		apiBody.Body = jokes[0]
	}
	return apiBody, err
}
