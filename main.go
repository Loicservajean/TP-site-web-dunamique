package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// Définition de la structure Article
type Article struct {
	ID        int
	Nom       string
	Prix      float64
	ImageUrl  string
	Reduction float64
}

// Liste des articles disponibles
var articles = []Article{
	{ID: 1, Nom: "PALACE PULL A CAPUCHE UNISÉXE CHASSEUR", Prix: 195.00, Reduction: 150.00, ImageUrl: "/assets/img/products/16A.webp"},
	{ID: 2, Nom: "PALACE PULL A CAPUCHE MARINE", Prix: 195.00, ImageUrl: "/assets/img/products/18A.webp"},
	{ID: 3, Nom: "PALACE PULL CREW PASSEPORT NOIR", Prix: 195.00, ImageUrl: "/assets/img/products/19A.webp"},
	{ID: 4, Nom: "PALACE WASHED TERRY 1/4 PLACKET HOOD MOJITO", Prix: 195.00, ImageUrl: "/assets/img/products/21A.webp"},
	{ID: 5, Nom: "PALACE PANTALON BOSSY JEAN STONE", Prix: 195.00, ImageUrl: "/assets/img/products/22A.webp"},
	{ID: 6, Nom: "PALACE PANTALON CARGO GORES-TEK™ NOIR", Prix: 195.00, ImageUrl: "/assets/img/products/33B.webp"},
	{ID: 7, Nom: "PALACE PANTALON CARGO GORES-TEK™ SABLE", Prix: 195.00, ImageUrl: "/assets/img/products/34B.webp"},
}

func main() {
	// Chargement des templates HTML
	listeTemplate, errTemplate := template.ParseGlob("template/*.html")
	if errTemplate != nil {
		fmt.Println("Erreur de chargement des templates :", errTemplate.Error())
		os.Exit(1)
	}

	// Route principale : affichage des produits
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := listeTemplate.ExecuteTemplate(w, "index", articles)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Route secondaire : liste de langages
	http.HandleFunc("/temp/list", func(w http.ResponseWriter, r *http.Request) {
		languages := []string{"Go", "Python", "JavaScript"}
		err := listeTemplate.ExecuteTemplate(w, "list", languages)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Route conditionnelle : exemple booléen
	http.HandleFunc("/temp/cond", func(w http.ResponseWriter, r *http.Request) {
		data := struct{ IsConnected bool }{IsConnected: true}
		err := listeTemplate.ExecuteTemplate(w, "cond", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Serveur des fichiers statiques (CSS, images)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Lancement du serveur
	fmt.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
