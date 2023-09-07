package main

import (
	// "fmt"
	"log"
	"microservices/handlers"
	"net/http"
	t "time"

	// "io"
	"os"
)

func main() {
	l := log.New(os.Stderr, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gb)

	s := &http.Server{
		Addr:    ":9090",
        Handler: sm,
		IdleTimeout:  120 * t.Second,
		ReadTimeout:  1   *t.Second,
        WriteTimeout: 1   *t.Second,
	}

	s.ListenAndServe()

	// Listen for connections on all ip addresses (0.0.0.0)
	// port 9090
	log.Println("Starting Server")
	http.ListenAndServe(":9090", sm)
	
}

