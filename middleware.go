package main

import (
	"fmt"
	"net/http"
	"time"
)

// Logger middleware
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		stop := time.Now()
		fmt.Printf("url: %s time: %s\n", r.URL.Path, stop.Sub(start))
	})
}
