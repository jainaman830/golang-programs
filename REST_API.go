package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	product := &Product{
		ProductID:   "P1",
		ProductName: "Drone",
		Price:       10000,
	}
	json.NewEncoder(w).Encode(product)
}
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}
func startServer() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/product", getProduct)
	http.ListenAndServe(":8080", nil)
}
func main() {
	startServer()
}
