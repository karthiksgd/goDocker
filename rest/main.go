package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karthiksgd/docker-rest-go/controller"
)

func main() {

	fmt.Println("Running Rest Server")

	r := mux.NewRouter()

	r.HandleFunc("/api/customers", controller.GetCustomers).Methods("Get")

	http.ListenAndServe(":8081", r)
}
