package handlers

import (
	// "microservices/data"
	// "fmt"
	"log"
	"microservices/data"
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
	jsson.Marshal
}