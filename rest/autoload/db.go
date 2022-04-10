package autoload

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(godockerDB)/sky_assess")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
