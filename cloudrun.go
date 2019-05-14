package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "Run"
	}

	for k, v := range r.Header {
		log.Printf("%s:%v\n", k, v)
	}

	w.Header().Set("cache-control", "public, max-age=3600")

	fmt.Fprintf(w, "Hello %s! %+v\n", target, time.Now())
}

func handlerNocache(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "Run"
	}
	w.Header().Set("cache-control", "private")
	fmt.Fprintf(w, "Hello %s Nocache %+v!\n", target, time.Now())
}

func main() {
	log.Print("Hello world sample started.")
	http.HandleFunc("/hellorun/nocache", handlerNocache)
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
