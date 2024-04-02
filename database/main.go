package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	id           int
	name         string
	phone_number string
}

func showTable(db *sql.DB) {
	log.Println("Showing tables")
	fmt.Println("------------------------------")
	defer fmt.Println("------------------------------")
	rows, err := db.Query("SELECT * FROM phone_book")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.id, &user.name, &user.phone_number); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d:\t%s (%s)\n", user.id, user.name, user.phone_number)
	}
}

func updateTable(db *sql.DB, name string, phone_number string) {
	_, err := db.Exec("INSERT INTO phone_book (name, phone_number) VALUES (?, ?)", name, phone_number)
	if err != nil {
		log.Fatal(err)
	}
}

func getInputData() (name string, phone_number string) {
	fmt.Println("\nEnter name and phone number")
	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("Phone number: ")
	fmt.Scan(&phone_number)
	return
}

func main() {
	db, err := sql.Open("sqlite3", "gopractice.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Connected to the database")

	showTable(db)

	for {
		name, phone_number := getInputData()
		if name == "" || phone_number == "" {
			log.Println("Name or phone number is empty")
			continue
		}
		updateTable(db, name, phone_number)
		showTable(db)
	}
}
