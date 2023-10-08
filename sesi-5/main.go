package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var PORT = ":8182"

type User struct {
	Email           string
	Password        string
	Address         string
	Job             string
	ReasonToBeHappy string
}

var users = []User{
	{Email: "squarepants@mail.com", Password: "sbob***", Address: "Pinaple house, Bikini Bottom", Job: "Chef", ReasonToBeHappy: "Patrick"},
	{Email: "star@mail.com", Password: "uhmmm", Address: "Under the rock, Bikini Bottom", Job: "Jobless", ReasonToBeHappy: "Lying under the rock"},
	{Email: "tentacles@mail.com", Password: "iloveme", Address: "Squid Gallery, Bikini Bottom", Job: "Artist", ReasonToBeHappy: "Nothing"},
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
	switch r.Method {
	case http.MethodGet:
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
	case http.MethodPost:
		email, password := r.FormValue("email"), r.FormValue("password")
		var user User
		for _, u := range users {
			if u.Email == email && u.Password == password {
				user = u
				break
			}
		}
		if user.Email == "" {
			falselogin, err := template.ParseFiles("./falseuser.html")
			if err != nil {
				log.Println("Error parsing falseuser.html:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = falselogin.Execute(w, nil)
			if err != nil {
				log.Println("Error executing falseuser.html:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		profile, err := template.ParseFiles("./profile.html")
		if err != nil {
			log.Println("Error parsing profile.html:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = profile.Execute(w, user)
		if err != nil {
			log.Println("Error executing profile.html:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
