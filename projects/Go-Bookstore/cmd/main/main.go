package main

import (
	"log"
	"net/http"
    _"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Devisrisamidurai/Go-mysql/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
