package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type handler struct{}

func randomStatus() int {
	status := rand.Intn(600)
		switch status {
		case http.StatusBadGateway, http.StatusNotFound, http.StatusInternalServerError, http.StatusServiceUnavailable:
			return status

		default:
			return http.StatusOK
		}
}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	duration := time.Duration(rand.Intn(100000)) * time.Microsecond
	ctx, cancel := context.WithTimeout(r.Context(), duration)
	defer cancel()

	select {
	case <-ctx.Done():
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, duration)
	}
}

func main() {
	http.ListenAndServe(":1323", handler{})
}
