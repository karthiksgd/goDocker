package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karthiksgd/sky_assessment/controller"
	"github.com/karthiksgd/sky_assessment/rest"
)

func main() {

	fmt.Println("Function Main in Terminal")

	// Router
	rx := mux.NewRouter()

	// Routes
	rx.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("view/assets"))))
	rx.HandleFunc("/", controller.LoginController)
	rx.HandleFunc("/logout", controller.LogoutController)
	rx.HandleFunc("/signup", controller.SignupController)
	rx.HandleFunc("/dashboard", controller.DashboardController)
	rx.HandleFunc("/listcustomers", controller.AllCustomersController)

	rx.HandleFunc("/api/adverts", rest.RestCalls).Methods("Post")

	// Run Server
	http.ListenAndServe(":8080", rx)

}
