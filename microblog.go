package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db gorm.DB
)

func connectToDB() {
	var err error
	db, err = gorm.Open("sqlite3", "microblog.db")
	if err != nil {
		fmt.Println(err)
	}
	db.DB()
}

func migrate() {
	db.AutoMigrate(&Post{})
}

func setUpRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/posts", postsHandler).Methods("GET")
	r.HandleFunc("/posts", newPost).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", r)
}

func main() {
	connectToDB()
	//migrate()
	setUpRoutes()
	fmt.Println("Server listening on port 8000...")
	http.ListenAndServe(":8000", nil)
}
