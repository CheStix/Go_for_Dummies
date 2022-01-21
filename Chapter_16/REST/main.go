package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the REST API!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", home)

	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
