package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Users []User

func connectDB() {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@(db)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Success")
	}
}

func responseTest(w http.ResponseWriter, r *http.Request) {
	users := Users{}
	for i := 0; i < 5; i++ {
		title := "Hello"
		users = append(
			users,
			User{Id: i, Name: title})
	}
	fmt.Println("call h2")
	json.NewEncoder(w).Encode(users)
}

func main() {
	connectDB()
	http.HandleFunc("/", responseTest)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
