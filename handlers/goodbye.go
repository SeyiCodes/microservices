package handlers

import (
	"fmt"
	"net/http"
	"log"
	"io"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}


func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    g.l.Println("Goodbye World")


	d, err := io.ReadAll(r.Body)
	if err!= nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		return
	}


	fmt.Fprintf(rw, "Goodbye %s", d)
}
