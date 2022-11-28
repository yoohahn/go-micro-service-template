package service

import (
	"context"
	"fmt"
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

func (ls *LoggingService) GetRandomJoke(ctx context.Context) (joke *ApiBody, err error) {
	defer func(start time.Time) {
		fmt.Printf("status=%v err=%s time=%v", joke.Code, err, time.Since(start))
	}(time.Now())
	return ls.next.GetRandomJoke(ctx)
}
