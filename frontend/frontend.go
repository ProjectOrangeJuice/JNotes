package frontend

import (
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type page struct {
	Title string
	Notes []string
}

func init() {
	var err error
	templates, err = template.ParseFiles("frontend/pages/header.html",
		"frontend/pages/index.html", "frontend/pages/footer.html")

	if err != nil {
		log.Fatalf("Failed to load templates %s", err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Fatalf("Failed execute index handler %s", err)
	}
}
