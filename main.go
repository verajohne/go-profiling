package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"runtime"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-redis/redis/v8"
)

func main() {
	runtime.SetBlockProfileRate(1)
	ctx := context.Background()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/debug", middleware.Profiler())

	r.Get("/data", func(w http.ResponseWriter, r *http.Request) {
		rdb := NewRedisClient()
		val, err := rdb.Get(ctx, "jsonData").Result()
		if err != nil {
			panic(err)
		}

		w.Write([]byte(val))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		rdb := NewRedisClient()
		resp, _ := http.Get("https://jsonplaceholder.typicode.com/todos/1")
		body, _ := ioutil.ReadAll(resp.Body)
		err := rdb.Set(ctx, "jsonData", body, 0).Err()
		if err != nil {
			panic(err)
		}
		w.Write([]byte("Written to Redis"))
	})

	r.Get("/httpclient", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := http.Get("https://jsonplaceholder.typicode.com/todos/1")
		body, _ := ioutil.ReadAll(resp.Body)
		w.Write(body)
	})

	http.ListenAndServe(":3000", r)
}

// while true; do
// curl localhost:3000/httpclient
// done

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
