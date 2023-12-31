package handlers

import (
	// "fmt"
	"log"
	"microservices/product-api/data"
	"net/http"
)



type Products struct {
	l *log.Logger
}

func NewProducts(l*log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err!= nil {
        http.Error(rw, "Unable to marshall data", http.StatusInternalServerError)
    }
}