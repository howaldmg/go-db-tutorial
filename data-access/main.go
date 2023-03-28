package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "example",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	checkAndPrintErr(err)
	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")
	checkAndPrintErr(err)
	fmt.Printf("Albums found: %v\n", albums)

	album, err := albumById(2)
	checkAndPrintErr(err)
	fmt.Printf("Album found: %v\n", album)

}

func checkAndPrintErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
