package controller

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"math/big"
	"net/http"

	"github.com/karthiksgd/docker-app-go/model"
)

type SignupPg struct {
	PageTitle         string
	ValidationMessage string
	Customers         []model.CustUser
}

func init() {

	tpl_admin = template.Must(tpl_admin.ParseGlob("view/components/*.gohtml"))
	tpl_admin = template.Must(tpl_admin.ParseGlob("view/*.gohtml"))

}

func SignupController(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Customer Signup Controller")

	pg := SignupPg{"Customer Signup | Sky Assessment", "", nil}

	if r.Method == http.MethodPost {

		if r.FormValue("password") == r.FormValue("cpassword") {

			userSignedup := model.User{1, r.FormValue("email"), r.FormValue("fname"), r.FormValue("lname"), "customer", false, r.FormValue("email"), r.FormValue("password")}

			uid := model.AddUser(&userSignedup)

			customerid, _ := rand.Int(rand.Reader, big.NewInt(999999999999))

			var subscription int
			var optin int

			if r.FormValue("subscription") == "on" {
				subscription = 1

			} else {
				subscription = 0
			}

			if r.FormValue("optin") == "on" {
				optin = 1

			} else {
				optin = 0
			}

			customerSignedup := model.Customer{1, uid, customerid, subscription, optin}

			cid := model.AddCustomer(&customerSignedup)

			if cid != 0 {

				fmt.Println(cid)
				pg.ValidationMessage = "Successfully Created Customer Account"

			} else {

				fmt.Println(cid)
				pg.ValidationMessage = "Error Occurred while creating Customer Account"
			}

		} else {

			pg.ValidationMessage = "Mismatch Password and Confirm Password"
		}
	}

	tpl_admin.ExecuteTemplate(w, "customersignup.gohtml", pg)

}

func AllCustomersController(w http.ResponseWriter, r *http.Request) {

	fmt.Println("List all Customers Controller")

	customers := model.AllCustomersModel()

	pg := SignupPg{"All Customers | Sky Assessment", "", customers}

	fmt.Printf("%+v", pg)

	tpl_admin.ExecuteTemplate(w, "listcustomers.gohtml", pg)
}
