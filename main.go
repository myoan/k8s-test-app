package main

import (
	"log"
	"net/http"
	"fmt"
	"os"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID int
	Name string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host     := os.Getenv("DB_HOST")
	dbname   := os.Getenv("DB_DBNAME")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, dbname))
	log.Println("Connected to DB!!")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT * FROM user") //
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)

		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(fmt.Sprintf("%d: %s\n", user.ID, user.Name)))
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
