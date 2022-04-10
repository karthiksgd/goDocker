package autoload

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() *sql.DB {

	os.Setenv("MYSQL_ROOT_USER", "root")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_ROOT_PASSWORD", "root")
	os.Setenv("MYSQL_DATABASE", "sky_assess")
	os.Setenv("MYSQL_HOST", "localhost")

	// db, err := sql.Open("mysql", "${MYSQL_ROOT_USER}:${MYSQL_ROOT_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}")
	db, err := sql.Open("mysql", "root:root@tcp(godockerDB)/sky_assess")

	if err != nil {
		log.Fatal(err)
	}
	return db
}
