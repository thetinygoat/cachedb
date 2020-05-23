package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// constants

// RouterConfig holds router config
type RouterConfig struct {
	router   *chi.Mux
	CacheRef GlobalCacheRef
}

// Response struct
type Response struct {
	Data  string `json:"data"`
	Error bool   `json:"error"`
}

var (
	serverPort string
)

func setConfigVars() {
	if len(os.Getenv("CACHEDB_PORT")) == 0 {
		serverPort = DefaultPort
	} else {
		serverPort = os.Getenv("CACHEDB_PORT")
	}
}

func (R *RouterConfig) initializeRouter(CacheRef GlobalCacheRef) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	R.router = r
	R.CacheRef = CacheRef
}

func (R *RouterConfig) registerRoutes() {
	r := R.router

	r.Get("/get", func(w http.ResponseWriter, r *http.Request) {

		key := r.URL.Query().Get("key")

		value, err := R.CacheRef.get(key)
		res := Response{}
		if err != nil {
			res.Data = err.Error()
			res.Error = true

			respondWithJSON(w, r, res)
		} else {
			res.Data = value
			res.Error = false

			respondWithJSON(w, r, res)
		}

	})
	r.Get("/set", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")
		ttlRaw := r.URL.Query().Get("ttl")

		ttl, err := strconv.Atoi(ttlRaw)

		res := Response{}
		if err != nil {
			res.Data = MalformedParams
			res.Error = true

			respondWithJSON(w, r, res)
		} else {
			response := R.CacheRef.set(key, value, ttl)
			res.Data = response
			res.Error = false

			respondWithJSON(w, r, res)
		}

	})
	r.Get("/del", func(w http.ResponseWriter, r *http.Request) {

		key := r.URL.Query().Get("key")
		response, err := R.CacheRef.del(key)
		res := Response{}
		if err != nil {
			res.Data = err.Error()
			res.Error = true

			respondWithJSON(w, r, res)
		} else {
			res.Data = response
			res.Error = false

			respondWithJSON(w, r, res)
		}

	})
	r.Get("/flush", func(w http.ResponseWriter, r *http.Request) {
		response := R.CacheRef.flush()
		res := Response{}
		res.Data = response
		res.Error = false

		respondWithJSON(w, r, res)

	})

	setConfigVars()
	log.Printf("Using port: %v", serverPort)
	srv := &http.Server{
		Addr:    ":" + serverPort,
		Handler: r,
	}

	log.Fatal(srv.ListenAndServe())
}
