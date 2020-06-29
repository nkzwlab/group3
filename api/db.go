package main

import (
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

/* User is structure to get information from db */
type User struct {
	Id        int    `json:"id"`
	LoginName string `json:"login_name"`
}

/* getUser gets an user with specified id */
func getUser(user_id int) (User, error) {
	// get a row
	row := db.QueryRow(`SELECT * FROM users WHERE id = ? LIMIT 1;`, user_id)

	user := User{}

	// scan data
	err := row.Scan(&user.Id, &user.LoginName)

	return user, err
}

/* getUserWithLoginName gets an user with specified login name */
func getUserWithLoginName(login_name string) (User, error) {
	// get a row
	row := db.QueryRow(`SELECT * FROM users WHERE login_name = ? LIMIT 1;`, login_name)

	user := User{}

	// scan data
	err := row.Scan(&user.Id, &user.LoginName)

	log.Printf("get user %v", user)

	return user, err
}

/* createUser creates an user with specified login_name */
func createUser(login_name string) (User, error) {
	// exec query
	_, err := db.Query(`INSERT INTO users(login_name) VALUES(?);`, login_name)
	if err != nil {
		return User{}, err
	}

	return getUserWithLoginName(login_name)
}

/* init initalizes database */
func init() {
	var err error
	// open DB
	db, err = sql.Open("mysql", "root:PASSWORD@tcp(db:3306)/db")

	// if some error occured, exit this program
	if err != nil {
		log.Fatal("Error: failed to open DB")
	}
}
