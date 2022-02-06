package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

type DashboardPg struct {
	PageTitle         string
	ValidationMessage string
}

func init() {

	tpl_admin = template.Must(tpl_admin.ParseGlob("view/components/*.gohtml"))
	tpl_admin = template.Must(tpl_admin.ParseGlob("view/*.gohtml"))

}

func DashboardController(w http.ResponseWriter, r *http.Request) {

	session, _ := storeSession.Get(r, "sessionData")

	if session.Values["isLoggedin"] == true {

		fmt.Println("Dashboard Controller")

		pg := DashboardPg{"Dashboard | Sky Assessment", "Status Goes here"}

		err := tpl_admin.ExecuteTemplate(w, "dashboard.gohtml", pg)

		if err != nil {

			panic(err)
		}

	} else {

		fmt.Println("Dashboard Controller... Not Logged in")

		pg := LoginPg{"Login | Sky Assessment", "You are logged out ..."}

		tpl_admin.ExecuteTemplate(w, "login.gohtml", pg)
	}

}
