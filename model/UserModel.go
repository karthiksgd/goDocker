package model

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/karthiksgd/sky_assessment/autoload"
)

var db *sql.DB

func init() {
	db = autoload.DbConn()
}

type User struct {
	Id       int
	Email    string
	Fname    string
	Lname    string
	Utype    string
	Status   bool
	Uname    string
	Password string
}

func Auth(e, p *string) bool {

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	var isAuth bool

	err = db.QueryRow(`SELECT IF (COUNT(*), true, false) FROM users WHERE email = ? AND password = ?;`, e, p).Scan(&isAuth)
	if err != nil {
		log.Fatal(err)
	}

	return isAuth
}

func GetUser(e, p *string) User {

	row, err := db.Query(`SELECT * FROM users WHERE email = ? AND password = ?;`, e, p)

	if err != nil {
		log.Fatal(err)
	}

	var user User

	for row.Next() {
		row.Scan(&user.Id, &user.Email, &user.Fname, &user.Lname, &user.Utype, &user.Status, &user.Uname, "")
	}

	return user
}

func AddUser(u *User) int64 {

	stmt := fmt.Sprintf("INSERT INTO `users`(`email`, `fname`, `lname`, `type`, `status`, `uname`, `password`) VALUES ('%v','%v','%v','%v','%v','%v','%v');", u.Email, u.Fname, u.Lname, u.Utype, 0, u.Uname, u.Password)

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
