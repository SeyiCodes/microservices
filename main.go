package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	t "time"
	"syscall"
	// "fmt"
	// "io"
	"microservices/product-api/handlers"
)

func main() {

	// env.Parse()

	// Create logger
	l := log.New(os.Stderr, "product-api", log.LstdFlags)

	// Create handlers for  routes
	// hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)

	// Create a ServeMux (multiplexer) to handle different route
	sm := http.NewServeMux()
	sm.Handle("/", ph)
	sm.Handle("/goodbye", gb)


	// Create an HTTP server
	s := &http.Server{
		Addr:          ":9090",
        Handler:       sm,
		IdleTimeout:   120 * t.Second,
		ReadTimeout:   1   *t.Second,
        WriteTimeout:  1   *t.Second,
	}


	// Start the server in a goroutine
	go func() {
        log.Println("Listening on port 9090")
		err := s.ListenAndServe()
		if err!= nil {
            l.Fatal(err)
        }
    }()

	// Set up a signal channel to gracefully handle termination signals
	c := make(chan os.Signal, 1)
	// added a buffer of 1


	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// Use syscall instead of os.Kill
	

	sig := <-c
	l.Println("Recieved Terminate Signal, shutting down gracefully", sig)

	

	// Create a context with a timeout
	tc, cancel := context.WithTimeout(context.Background(), 30*t.Second)
	defer cancel()	
		
	if err:= s.Shutdown(tc); err != nil {
		l.Fatalf("Error during server shutdown: %v", err)
	}

	
	// log.Println("Starting Server")
	// http.ListenAndServe(":9090", sm)
	
}

