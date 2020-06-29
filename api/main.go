package main

import (
	"database/sql"
	_ "fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	db *sql.DB
)

/* main routine */
func main() {
	defer db.Close()

	router := httprouter.New()

	// routings
	router.GET("/user", GetUser)
	router.POST("/user/new", CreateUser)

	router.POST("/kadai/new", CreateKadai)
	//rouer.GET("/kadai", KadaiIndex)
	//rouer.PATCH("/kadai/update", UpdateKadai)
	//rouer.PATCH("/kadai/done", KadaiDone)

	// start server
	log.Printf("Start Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
