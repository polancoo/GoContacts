package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

	// Captura los valores para conexion a la BDD MySQL.
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "D_aleLike",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "go_contacts",
		AllowNativePasswords: true,
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Conectado a la BDD!")
}
