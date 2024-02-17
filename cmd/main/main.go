package main

import (
	"log"
	"net/http"

	_ "github.com/jinhzu/gorm/dialects/mysql"

	"github.com/gorilla/mux"
	"github.com/harshakp06/mysql/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
