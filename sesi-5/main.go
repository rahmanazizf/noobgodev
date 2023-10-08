package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var PORT = ":8182"

type User struct {
	Email    string
	Password string
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/login", login)
	log.Println("Listening on localhost", PORT)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, Mom!")
	if err != nil {
		log.Fatal(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.Method == http.MethodGet {
		tpl, err := template.ParseFiles("./template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Error parsing template:", err)
			return
		}
		err = tpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Error executing template:", err)
			return
		}
		return
	} else {
		email, password := r.FormValue("email"), r.FormValue("password")
		user := User{Email: email, Password: password}
		if email == "rahmanazizf@mail.com" && password == "razizf" {
			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println("Error encoding JSON:", err)
				return
			}
			return
		}
	}

	http.Error(w, "Invalid email or password", http.StatusUnauthorized)
}
