package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func register(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = db.QueryRow("SELECT username FROM users WHERE username = ?", user.Username).Scan(&user.Username)
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		return
	} else if err != sql.ErrNoRows {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Id = uuid.New()
	user.Password = string(hashed)
	_, err = db.Exec("INSERT INTO users (uuid,username,password) VALUES (?,?,?)", user.Id, user.Username, user.Password)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Print("New User created")
	w.WriteHeader(http.StatusCreated)
}
func login(w http.ResponseWriter, r *http.Request) {
	log.Print("New User")
	var login Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	var user User
	rows, err := db.Query("SELECT * FROM users WHERE username = ?", login.Username)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
}
