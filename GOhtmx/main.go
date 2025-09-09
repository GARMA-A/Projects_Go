package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"films": {
				{Title: "The GodFather", Director: "Franc ford"},
				{Title: "Blade Runner", Director: "Ridley scot"},
				{Title: "The Things", Director: "Jhon Carpenter"},
			},
		}
		templ.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s- %s</li>", title, director)
		templ, _ := template.New("t").Parse(htmlStr)
		templ.Execute(w, nil)
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
