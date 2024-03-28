package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {

	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Buy milk", Done: true},
				{Title: "Buy eggs", Done: false},
				{Title: "Buy bread", Done: true},
			},
		}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}
