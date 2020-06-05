package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// Logger middleware
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ua := r.Header.Get("User-Agent")
		next.ServeHTTP(w, r)
		stop := time.Now()
		fmt.Printf("%s %s %s %s\n", r.URL.Path, stop.Sub(start), ua, start)
	})
}

// Auth middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Header.Get("X-Cachedb-User")
		password := r.Header.Get("X-Cachedb-Password")
		if user == "" || password == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			return
		}
		if user != os.Getenv(CachedbUser) || password != os.Getenv(CachedbPassword) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
