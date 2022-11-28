package service

type RandomJoke struct {
	Joke string `json:"joke"`
}

type EmptyStruct struct{}

type ApiBody struct {
	Code      int        `json:"code"`
	Body      RandomJoke `json:"body"`
	Errors    []string   `json:"errors"`
	HasErrors bool       `json:"hasErrors"`
}
