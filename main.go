package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type Post struct {
	Id    uuid.UUID `json:"uuid"`
	Owner uuid.UUID `json:"owner"`
	Title string    `json:"title"`
	Body  string    `json:"body"`
}
type User struct {
	Id       uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
type Session struct {
	Id   uuid.UUID `json:"uuid"`
	User uuid.UUID `json:"user"`
}

var db *sql.DB

func main() {
	Connect()
	http.HandleFunc("/posts", getPosts)
	http.HandleFunc("/auth/login", login)
	http.HandleFunc("/auth/register", register)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello World")
}
func getPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Print("Error querying database ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Id, &post.Title, &post.Body)
		if err != nil {
			log.Print("Error scanning values ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}
	enc := json.NewEncoder(w)
	if enc := enc.Encode(posts); enc != nil {
		log.Print("Error encoding JSON ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
