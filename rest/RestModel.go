package rest

import (
	"database/sql"
	"log"

	"github.com/karthiksgd/sky_assessment/autoload"
)

var db *sql.DB

func init() {
	db = autoload.DbConn()
}

type SkyCustomers struct {
	SkyId        int64 `json:"customerId"`
	Subscription bool  `json:"subscription"`
	Optin        bool  `json:"optin"`
}

func GetCustomerData() []SkyCustomers {

	rows, err := db.Query("SELECT customer_id, subscription, ads_optin FROM `customers`")
	if err != nil {
		log.Fatal(err)
	}

	var data = SkyCustomers{}
	var cust = []SkyCustomers{}

	for rows.Next() {

		rows.Scan(&data.SkyId, &data.Subscription, &data.Optin)
		cust = append(cust, data)
	}

	return cust
}
