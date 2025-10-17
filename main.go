package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Article struct {
	ID       int
	Nom      string
	Prix     float64
	ImageUrl string
}

func main() {
	listeTemplate, errTemplate := template.ParseGlob("template/*.html")
	if errTemplate != nil {
		fmt.Println(errTemplate.Error())
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		articles := []Article{
			{ID: 1, Nom: "PALACE PULL A CAPUCHE UNISSEX CHASSEUR", Prix: 195, ImageUrl: "/assets/img/products/1.jpg"},
			{ID: 2, Nom: "PALACE PULL A CAPUCHE MARINE", Prix: 195, ImageUrl: "/assets/img/products/2.jpg"},
			{ID: 3, Nom: "PALACE PULL CREW PASSAGER NOIR", Prix: 195, ImageUrl: "/assets/img/products/3.jpg"},
		}
		listeTemplate.ExecuteTemplate(w, "index", articles)
	})

	http.HandleFunc("/temp/list", func(w http.ResponseWriter, r *http.Request) {
		languages := []string{"Go", "Python", "JavaScript"}
		listeTemplate.ExecuteTemplate(w, "list", languages)
	})

	http.HandleFunc("/temp/cond", func(w http.ResponseWriter, r *http.Request) {
		data := struct{ IsConnected bool }{IsConnected: true}
		listeTemplate.ExecuteTemplate(w, "cond", data)
	})

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8080", nil)
}
