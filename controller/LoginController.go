package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/karthiksgd/sky_assessment/model"
)

var tpl_admin *template.Template

type LoginPg struct {
	PageTitle         string
	ValidationMessage string
}

var (
	key          = []byte("sky_assess")
	storeSession = sessions.NewCookieStore(key)
)

func init() {

	tpl_admin = template.Must(tpl_admin.ParseGlob("view/*.gohtml"))
}

func LoginController(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Login Controller")

	pg := LoginPg{"Login | Sky Assessment", ""}

	session, err := storeSession.Get(r, "sessionData")

	if err != nil {
		panic(err)
	}

	if r.Method == http.MethodPost {

		// Get Form Data
		email := r.FormValue("email")
		password := r.FormValue("password")

		fmt.Println(email, password)

		// Authenticate User
		auth := model.Auth(&email, &password)

		if auth {

			// Get User Data
			user := model.GetUser(&email, &password)
			// fmt.Printf("%+v\n", user)

			// Create Session
			if user.Utype == "admin" && user.Status == true {

				session.Values["uId"] = user.Id
				session.Values["uEmail"] = user.Email
				session.Values["uType"] = user.Utype
				session.Values["uName"] = user.Uname

				err := session.Save(r, w)
				if err != nil {
					panic(err)
				}
				fmt.Printf("%#v\n", session)

				// http.Redirect(w, r, "/dashboard", 303)

			}

			// Else Condition Continues ...

		} else {

			pg = LoginPg{"Login | Sky Assessment", "Invalid Credentials"}

		}
	}

	tpl_admin.ExecuteTemplate(w, "login.gohtml", pg)
}
