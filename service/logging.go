package service

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (ls *LoggingService) GetRandomJoke(ctx context.Context, r *http.Request) (joke *ApiBody, err error) {
	defer func(start time.Time) {
		// Send log to logging service
		fmt.Printf("\nstatus=%v path=%s err=%s time=%v\n", joke.Code, r.URL.Path, err, time.Since(start))
	}(time.Now())
	return ls.next.GetRandomJoke(ctx, r)
}
