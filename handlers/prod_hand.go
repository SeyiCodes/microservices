package handlers

import (
	// "fmt"
	"log"
	"microservices/data"
	"net/http"
	"encoding/json"
)



type Products struct {
	l *log.Logger
}

func NewProducts(l*log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err!= nil {
        http.Error(rw, "Unable to marshall data", http.StatusInternalServerError)
    }

	rw.Write(d)
}