package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karthiksgd/docker-app-go/controller"
	// "github.com/karthiksgd/sky_assessment/rest"
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

	// rx.HandleFunc("/api/customers", rest.GetCustomers).Methods("Get")
	// rx.HandleFunc("/api/custdata", rest.CustomerData).Methods("Post")
	// rx.HandleFunc("/api/checkcustomer", rest.CheckCustomer).Methods("Post")

	// Run Server
	http.ListenAndServe(":8080", rx)

}
