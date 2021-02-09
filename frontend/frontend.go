package frontend

import (
	"log"
	"net/http"
	"text/template"

	"github.com/projectorangejuice/jnotes/sql"
)

var templates *template.Template

type page struct {
	Title string
	Notes []string
}

func init() {
	var err error
	templates, err = template.ParseFiles("frontend/pages/header.t", "frontend/pages/index.t", "frontend/pages/side.t")
	if err != nil {
		log.Fatalf("Failed to load templates %s", err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	p := page{Title: "okay", Notes: sql.GetNotes()}
	err := templates.ExecuteTemplate(w, "index", p)
	if err != nil {
		log.Fatalf("Failed execute index handler %s", err)
	}
}
