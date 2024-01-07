package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradeepjaju1/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
