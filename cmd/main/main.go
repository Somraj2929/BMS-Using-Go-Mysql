package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"github.com/somraj/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server running on http://localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}