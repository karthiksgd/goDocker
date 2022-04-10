package model

import (
	"fmt"
	"log"
	"math/big"

	"github.com/karthiksgd/docker-app-go/autoload"
)

func init() {
	db = autoload.DbConn()
}

type Customer struct {
	Id           int
	UserId       int64
	CustomerId   *big.Int
	Subscription int
	Optin        int
}

type CustUser struct {
	Id           int
	UserId       int64
	Fname        string
	Lname        string
	Email        string
	CustomerId   int64
	Subscription int
	Optin        int
}

func AddCustomer(c *Customer) int64 {

	stmt := fmt.Sprintf("INSERT INTO `customers`(`user_id`, `customer_id`, `subscription`, `ads_optin`) VALUES ('%v','%v','%v','%v');", c.UserId, c.CustomerId, c.Subscription, c.Optin)

	query, err := db.Prepare(fmt.Sprintf("%v", stmt))

	if err != nil {
		log.Fatal(err)
	}

	stat, err := query.Exec()
	if err != nil {
		log.Fatal(err)
	}

	id, err := stat.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	return id
}

func AllCustomersModel() []CustUser {

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM `customers` INNER JOIN `users`ON users.id = customers.user_id;")

	if err != nil {
		log.Fatal(err)
	}

	var cust = CustUser{}
	var custs = []CustUser{}

	for rows.Next() {

		rows.Scan(&cust.Id, &cust.UserId, &cust.CustomerId, &cust.Subscription, &cust.Optin, &cust.Id, &cust.Email, &cust.Fname, &cust.Lname, "", "", "", "")
		fmt.Println("Customer ", cust)
		custs = append(custs, cust)
	}

	return custs
}
