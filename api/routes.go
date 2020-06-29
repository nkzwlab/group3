package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

/* GetUser returns user info */
func GetUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logUrl(r)

	// get form value
	user_id, _ := strconv.Atoi(r.FormValue("user_id"))
	login_name := r.FormValue("login_name")

	user := User{}
	var err error

	// get user
	if user_id != 0 {
		user, err = getUser(user_id)
	} else if login_name != "" {
		user, err = getUserWithLoginName(login_name)
	} else {
		respondError(w, "invalid form value")
		return
	}

	if err != nil {
		respondError(w, "failed to get user")
		return
	}

	// return json
	respondJson(w, user)
}

/* CreateUser creates user info */
func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logUrl(r)

	// get form values
	login_name := r.FormValue("login_name")

	if login_name == "" {
		respondError(w, "invalid form value. login_name can't be empty")
		return
	}

	log.Println("login name:", login_name)

	// create user
	user, err := createUser(login_name)
	if err != nil {
		respondError(w, "failed to create user")
		return
	}

	respondJson(w, user)
}

type ErrorJson struct {
	Error string `json:"error"`
}

func respondError(w http.ResponseWriter, err interface{}) {
	data := ErrorJson{err.(string)}

	respondJson(w, data)
}

func respondJson(w http.ResponseWriter, data interface{}) {
	log.Printf("response: %v", data)

	// return json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func logUrl(r *http.Request) {
	log.Printf("access: %v", r.URL)
}
