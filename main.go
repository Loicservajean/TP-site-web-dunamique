package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type Article struct {
	ID          int
	Nom         string
	Prix        float64
	ImageUrl    string
	Reduction   float64
	Description string
}

// Liste des articles
var articles = []Article{
	{ID: 1, Nom: "PALACE PULL A CAPUCHE UNISÉXE CHASSEUR", Prix: 195.00, Reduction: 150.00, ImageUrl: "/assets/img/products/16A.webp", Description: "Conçu pour les puristes du style urbain, ce hoodie incarne l’équilibre parfait entre confort thermique et esthétique brute. Le marine profond contraste avec le logo Palace pour un rendu sobre mais affirmé"},
	{ID: 2, Nom: "PALACE PULL A CAPUCHE MARINE", Prix: 195.00, ImageUrl: "/assets/img/products/18A.webp", Description: "Alliant fonctionnalité et style, ce hoodie est conçu pour offrir une chaleur optimale grâce à son tissu épais, tout en arborant un design épuré avec le logo Palace brodé, parfait pour les amateurs de mode urbaine."},
	{ID: 3, Nom: "PALACE PULL CREW PASSEPORT NOIR", Prix: 195.00, ImageUrl: "/assets/img/products/19A.webp", Description: "Ce sweat crew incarne l’essence du style urbain avec son design minimaliste et son logo Palace brodé. Confectionné en coton épais, il offre un confort supérieur et une durabilité, idéal pour les journées fraîches en ville."},
	{ID: 4, Nom: "PALACE WASHED TERRY 1/4 PLACKET HOOD MOJITO", Prix: 195.00, ImageUrl: "/assets/img/products/21A.webp", Description: "Ce hoodie 1/4 placket en tissu éponge lavé offre un confort exceptionnel et un style décontracté. Sa teinte mojito rafraîchissante et son design unique en font une pièce incontournable pour ceux qui recherchent à la fois confort et originalité."},
	{ID: 5, Nom: "PALACE PANTALON BOSSY JEAN STONE", Prix: 195.00, ImageUrl: "/assets/img/products/22A.webp", Description: "Ce jean Bossy en stone wash offre un look décontracté avec une touche vintage. Sa coupe droite et son tissu en denim robuste garantissent confort et durabilité, tandis que les détails subtils ajoutent une note de style unique."},
	{ID: 6, Nom: "PALACE PANTALON CARGO GORES-TEK™ NOIR", Prix: 195.00, ImageUrl: "/assets/img/products/33B.webp", Description: "Alliant fonctionnalité et style, ce pantalon cargo Gores-Tek™ noir est conçu pour résister aux éléments tout en offrant un confort optimal. Ses multiples poches pratiques et sa coupe moderne en font un choix idéal pour les aventures urbaines ou en plein air."},
	{ID: 7, Nom: "PALACE PANTALON CARGO GORES-TEK™ SABLE", Prix: 195.00, ImageUrl: "/assets/img/products/34B.webp", Description: "Ce pantalon cargo Gores-Tek™ en sable combine durabilité et style avec son tissu résistant aux intempéries et sa coupe moderne. Parfait pour les explorateurs urbains, il offre de nombreuses poches pratiques pour un rangement facile."},
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

	// Route article : affichage des détails d'un produit
	http.HandleFunc("/article", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		for _, a := range articles {
			if fmt.Sprintf("%d", a.ID) == id {
				listeTemplate.ExecuteTemplate(w, "article", a)
				return
			}
		}
		http.Error(w, "Article non trouvé", http.StatusNotFound)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			listeTemplate.ExecuteTemplate(w, "add", nil)
			return
		}

		if r.Method == http.MethodPost {
			nom := r.FormValue("nom")
			prixStr := r.FormValue("prix")
			reductionStr := r.FormValue("reduction")
			image := r.FormValue("image")
			description := r.FormValue("description")

			prix, _ := strconv.ParseFloat(prixStr, 64)
			reduction, _ := strconv.ParseFloat(reductionStr, 64)

			newID := len(articles) + 1
			article := Article{
				ID:          newID,
				Nom:         nom,
				Prix:        prix,
				Reduction:   reduction,
				ImageUrl:    image,
				Description: description,
			}
			articles = append(articles, article)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	})

	// Serveur des fichiers statiques (CSS, images)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Lancement du serveur go
	fmt.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
