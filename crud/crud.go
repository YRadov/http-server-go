package crud

import (
	"net/http"
	"http_server/model"
	"http_server/utils/render"
	"http_server/utils/validation"
)


func ViewHandler(w http.ResponseWriter, r *http.Request) {

	title, err := validation.GetTitle(w, r)

	if err != nil {
		return
	}

	p, err := model.LoadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	render.Template(w, "view", p)

} //ViewHandler

func EditHandler(w http.ResponseWriter, r *http.Request) {

	title, err := validation.GetTitle(w, r)

	if err != nil {
		return
	}

	p, err := model.LoadPage(title)

	if err != nil {
		p = &model.Page{Title: title}
	}

	render.Template(w, "edit", p)

} //EditHandler

func SaveHandler(w http.ResponseWriter, r *http.Request) {

	title, err := validation.GetTitle(w, r)

	if err != nil {
		return
	}

	body := r.FormValue("body")

	p := &model.Page{Title: title, Body: []byte(body)}

	err = p.Save()

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)

} //SaveHandler
