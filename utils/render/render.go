package render

import (
	"net/http"
	"http_server/model"
	"html/template"
)

/** кешируем шаблоны */
var templates = template.Must(template.ParseFiles("templates/edit.html",
	"templates/view.html"))

func Template(w http.ResponseWriter, tmpl string, p *model.Page) {

	/** используем шаблон напрямую */
	//t, err := template.ParseFiles("templates/" + tmpl + ".html")
	//
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//err = t.Execute(w, p)
	//
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}

	/** достаем нужный шаблон из кеша */
	err := templates.ExecuteTemplate(w, tmpl+".html", p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

} //Template
