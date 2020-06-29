package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

/* Root just returns user info */
func Root(w http.ResponseWriter, r *http.Request) {
	log.Printf("access: %v", r.URL)

	// get form value
	login_name := r.FormValue("login_name")

	// get user
	info, err := getUser(login_name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to get user\n")
		return
	}

	// return json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

/* CreateUser creates user info */
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// CreateUser accepts only POST requests
	if r.Method != "POST" {
		Root(w, r)
		return
	}

	log.Printf("access: %v", r.URL)

	// get form values
	name := r.FormValue("name")
	kg := r.FormValue("kg")
	login_name := r.FormValue("login_name")

	// all fields mustn't be empty
	if name == "" || kg == "" || login_name == "" {
		fmt.Fprintf(w, "invalid form value\n")
		fmt.Fprintf(w, "you need to post name, kg and login name\n")
		return
	}

	log.Println("user info:", name, kg, login_name)

	// create user
	_, err := db.Query("INSERT INTO info(name, kg, login_name) VALUES (?, ?, ?)", name, kg, login_name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to create user\n")
		return
	}

	// get user
	info, err := getUser(login_name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to get user\n")
		return
	}

	// return json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

/* Intro is structure to get information from db */
type Intro struct {
	Id        int    `json:"-"`
	Name      string `json:"name"`
	Kg        string `json:"kg"`
	LoginName string `json:"login_name"`
}

/* getUser gets an user based on login name */
func getUser(login_name string) (Intro, error) {
	// get a row
	row := db.QueryRow(`SELECT * FROM info WHERE login_name LIKE ? LIMIT 1;`, "%"+login_name+"%")

	info := Intro{}

	// scan data
	err := row.Scan(&info.Id, &info.Name, &info.Kg, &info.LoginName)

	return info, err
}

/* main routine */
func main() {
	defer db.Close()

	// routings
	http.HandleFunc("/", Root)
	http.HandleFunc("/user/new", CreateUser)

	// start server
	log.Printf("Start Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
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
