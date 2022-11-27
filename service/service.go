package service

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetRandomJoke(context.Context) (*RandomJoke, error)
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

func (rjs *RandomJokeService) GetRandomJoke(ctx context.Context) (*RandomJoke, error) {
	resp, err := http.Get(rjs.url)
	resp.Header.Add("X-Api-Key", rjs.apiKey)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	joke := &RandomJoke{}
	if err := json.NewDecoder(resp.Body).Decode(joke); err != nil {
		return nil, err
	}
	return joke, err
}

// urlData := UrlData{Limit: 1, Version: "v1", Key: GetApiKey()}
// 	tmpl, err := template.New("test").Parse("https://api.api-ninjas.com/{{.Version}}/dadjokes?limit={{.Limit}}")
