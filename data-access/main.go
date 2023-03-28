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
		User:                 "root",    //tutorial loads from os.Getenv()
		Passwd:               "example", //tutorial loads from os.Getenv()
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	checkAndPrintErr(err)

	err = db.Ping()
	checkAndPrintErr(err)
	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")
	checkAndPrintErr(err)
	fmt.Printf("Albums found: %v\n", albums)

	album, err := albumById(2)
	checkAndPrintErr(err)
	fmt.Printf("Album found: %v\n", album)

	albId, err := addAlbum(Album{
		Title:  "The Modern Sound of Better Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	checkAndPrintErr(err)
	fmt.Printf("ID of added album: %v\n", albId)
}

func checkAndPrintErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
