package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string
	Lastname  string
	Age       int
}

func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, "%s %s is %d years old", user.Firstname, user.Lastname, user.Age)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "Peter",
			Lastname:  "Griffin",
			Age:       50,
		}

		json.NewEncoder(w).Encode(peter)
	})

	http.ListenAndServe(":8080", nil)
}
