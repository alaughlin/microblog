package main

import (
	"encoding/json"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/root.html")
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	var posts []Post
	db.Order("created_at desc").Find(&posts)
	res, _ := json.Marshal(posts)
	w.Write([]byte(res))
}

func newPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	post := Post{Title: r.Form.Get("title"), Content: r.Form.Get("content")}
	db.Create(&post)
	res, _ := json.Marshal(post)
	w.Write([]byte(res))
}
