package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Password string
}

func main() {
	const (
		host     = "127.0.0.1"
		port     = 5432
		user     = "user"
		password = "password"
		dbname   = "db_schema"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening example: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the example: %v", err)
	}

	fmt.Println("Successfully connected to the example!")
	rows, err := db.Query(`SELECT id, username, password FROM users limit 10`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Username, &u.Password)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Retrieved users:", users)
}
