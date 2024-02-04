package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
}

var Allproducts []Product

func GetProduct(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id, ok := urlParams["pid"]
	if ok {
		for _, product := range Allproducts {
			if product.ProductID == id {
				json.NewEncoder(w).Encode(product)
			}
		}
	} else {
		json.NewEncoder(w).Encode(Allproducts)
	}

}
func AddProduct(w http.ResponseWriter, r *http.Request) {
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error in body read : %f", err)
	}
	newproduct := Product{}
	err = json.Unmarshal(rBody, &newproduct)
	if err != nil {
		log.Fatalf("Invalid Input : %f", err)
	}
	productFound := false
	for _, product := range Allproducts {
		if product.ProductID == newproduct.ProductID {
			productFound = true
			break
		}
	}
	if productFound {
		json.NewEncoder(w).Encode("Product id already exists.")
	} else {
		Allproducts = append(Allproducts, newproduct)
		json.NewEncoder(w).Encode("Success")
	}

}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	rBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error in body read : %f", err)
	}
	newproduct := Product{}
	err = json.Unmarshal(rBody, &newproduct)
	if err != nil {
		log.Fatalf("Invalid Input : %f", err)
	}
	if newproduct.ProductID == "" {
		json.NewEncoder(w).Encode("Provide Product details.")
	} else {
		productFound := false
		for i, product := range Allproducts {
			if product.ProductID == newproduct.ProductID {
				productFound = true
				Allproducts[i].ProductName = newproduct.ProductName
				Allproducts[i].Price = newproduct.Price
				break
			}
		}
		if productFound {
			json.NewEncoder(w).Encode("Product Updated.")
		} else {
			json.NewEncoder(w).Encode("Provide Valid Product!")
			Allproducts = append(Allproducts, newproduct)
			json.NewEncoder(w).Encode("Success")
		}
	}
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id, ok := urlParams["pid"]
	if ok {
		ProductFound := false
		for i, product := range Allproducts {
			if product.ProductID == id {
				ProductFound = true
				Allproducts = append(Allproducts[:i], Allproducts[i+1:]...)
			}
		}
		if ProductFound {
			json.NewEncoder(w).Encode("Product deleted.")
		} else {
			json.NewEncoder(w).Encode("Provide valid product id!")
		}
	} else {
		json.NewEncoder(w).Encode("Provide product id!")
	}
}
func ProductHomepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}
func StartGorillaMuxServer() {
	server := mux.NewRouter().StrictSlash(true)
	server.HandleFunc("/", ProductHomepage)
	server.HandleFunc("/product", GetProduct)
	server.HandleFunc("/product/{pid}", GetProduct)
	server.HandleFunc("/addproduct", AddProduct).Methods("POST")
	server.HandleFunc("/updateproduct", UpdateProduct).Methods("PUT")
	server.HandleFunc("/deleteproduct/{pid}", DeleteProduct).Methods("DELETE")
	http.ListenAndServe(":8080", server)
}
func main() {
	Allproducts = []Product{
		Product{
			ProductID:   "P1",
			ProductName: "Drone",
			Price:       10000,
		},
		Product{
			ProductID:   "P2",
			ProductName: "Mobile",
			Price:       5000,
		},
	}
	StartGorillaMuxServer()
}
