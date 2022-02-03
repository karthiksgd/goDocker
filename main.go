package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karthiksgd/sky_assessment/rest"
)

func main() {

	fmt.Println("Function Main in Terminal")

	// Router
	rx := mux.NewRouter()

	// Routes
	rx.HandleFunc("/", rest.RestCalls).Methods("Post")

	http.ListenAndServe(":8080", rx)

}
