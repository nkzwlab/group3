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

/* Kadai is structure to get kadai information from db */
type Kadai struct {
	Id      int    `json:"id"`
	UserId  int    `json:"-"`
	User    User   `json:"user"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Draft   string `json:"draft"`
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

/* getKadai gets a kadai with specified id */
func getKadai(kadai_id int) (Kadai, error) {
	row := db.QueryRow(`SELECT * FROM kadai WHERE id = ? LIMIT 1;`, kadai_id)

	kadai := Kadai{}
	err := row.Scan(&kadai.Id, &kadai.UserId, &kadai.Title, &kadai.Content, &kadai.Draft)
	if err != nil {
		log.Println("error:", err)
		return Kadai{}, err
	}

	kadai.User, err = getUser(kadai.UserId)
	if err != nil {
		log.Println("error:", err)
		return Kadai{}, err
	}

	return kadai, nil
}

/* createKadai creates a kadai with specified params */
func createKadai(user_id int, title, content, draft string) (Kadai, error) {
	_ = db.QueryRow(`INSERT INTO kadai(user_id, title, content, draft) VALUES (?, ?, ?, ?)`, user_id, title, content, draft)

	row := db.QueryRow(`SELECT id FROM kadai WHERE user_id = ? AND title = ? AND content = ? AND draft = ? LIMIT 1;`, user_id, title, content, draft)

	var kadai_id int
	err := row.Scan(&kadai_id)
	if err != nil {
		return Kadai{}, err
	}
	log.Printf("created kadai %v", kadai_id)

	return getKadai(kadai_id)
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
