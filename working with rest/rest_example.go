package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint Hit!")
	fmt.Println("Welcome to the Home Page")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", router))
}