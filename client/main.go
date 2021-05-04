package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/progfay/kataribe-playgound/client/request"
	"github.com/progfay/kataribe-playgound/client/thread"
	"golang.org/x/sync/errgroup"
)

const (
	concurrency = 4
	duration    = 10 * time.Second
	host        = "http://localhost:8080"
)

var requestLines = []request.RequestLine{
	request.NewRequestLine("GET", host, "/"),
	request.NewRequestLine("GET", host, "/users"),
	request.NewRequestLine("GET", host, "/users/:id"),
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)

	for i := 0; i < concurrency; i++ {
		eg.Go(func() error {
			t := thread.New()
			return t.Run(ctx, requestLines)
		})
	}

	if err := eg.Wait(); !errors.Is(err, context.DeadlineExceeded) {
		log.Printf("%+v\n", err)
	}
}
