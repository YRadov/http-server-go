package render

import (
	"net/http"
	"http_server/model"
	"html/template"
)


func Template(w http.ResponseWriter, tmpl string, p *model.Page) {
	t, _ := template.ParseFiles("templates/" + tmpl + ".html")
	t.Execute(w, p)
}
