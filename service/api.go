package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiServer struct {
	svc Service
}

func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func (s *ApiServer) Start(addrs string) error {
	fmt.Println("Server: http://localhost" + addrs)
	http.HandleFunc("/", s.handleGetRandomJoke)
	return http.ListenAndServe(addrs, nil)
}

func (s *ApiServer) handleGetRandomJoke(w http.ResponseWriter, r *http.Request) {
	apiBody, err := s.svc.GetRandomJoke(context.Background(), r)
	if err != nil {
		// Send error to logging service
		fmt.Printf("Err=%+v\n", err)
	}
	writeBodyAsJson(w, apiBody)
}

func writeBodyAsJson(w http.ResponseWriter, body *ApiBody) error {
	w.WriteHeader(body.Code)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(body)
}
