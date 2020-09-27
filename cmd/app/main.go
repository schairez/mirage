package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// func Connect(host string, port int, user string, password string, dbname string) (*sql.DB, error) {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

type User struct {
	ID       int
	Email    string
	Password string
}

func main() {
	fmt.Println("yo")
	fmt.Println(os.Getenv("POSTGRES_HOST"))
	// Create DB pool
	// cfg :=
	// db, err := database.New()
	// db, err := sql.Open("postgres", DB_DSN)
	// if err != nil {
	// 	log.Fatal("Failed to open a DB connection: ", err)
	// }
	// defer db.Close()

	// // Create an empty user and make the sql query (using $1 for the parameter)
	// var myUser User
	// userSql := "SELECT id, email, password FROM users WHERE id = $1"

	// err = db.QueryRow(userSql, 1).Scan(&myUser.ID, &myUser.Email, &myUser.Password)
	// if err != nil {
	// 	log.Fatal("Failed to execute query: ", err)
	// }

	// fmt.Printf("Hi %s, welcome back!\n", myUser.Email)
}
