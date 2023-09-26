package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Product struct {
	Id string
	Name string
	Quantity int
	Price float64
} 

var Products []Product;

func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: homepage")
	fmt.Fprintf(w, "Welcome to the homepage")
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: returnAllProducts")
	json.NewEncoder(w).Encode(Products)
}

func returnProductById(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	vars := mux.Vars(r)
	key := vars["id"]
	for _, product := range Products {
		if product.Id == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/products", returnAllProducts)
	myRouter.HandleFunc("/product/{id}", returnProductById)
	myRouter.HandleFunc("/", homepage)
	http.ListenAndServe("localhost:10000", myRouter)
}

func main() {

	Products = []Product{
		{Id: "1", Name: "Chair", Quantity: 100, Price: 100.00},
		{Id: "2", Name: "Desk", Quantity: 200, Price: 200.00},
	}

	handleRequests();
}