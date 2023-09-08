package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye World")

	// Use io.ReadAll directly in the response writer
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	// Set the response content type and status code
	rw.Header().Set("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusOK)

	// Write the response message
	fmt.Fprintf(rw, "Goodbye %s", d)
}
