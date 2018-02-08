package crud

import (
	"net/http"
	"http_server/model"
	"http_server/utils/render"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/view/"):]

	p, err := model.LoadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)

		return
	}

	//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

	render.Template(w, "view", p)

} //ViewHandler

func EditHandler(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/edit/"):]

	p, err := model.LoadPage(title)

	if err != nil {
		p = &model.Page{Title: title}
	}

	//fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	//	"<form action=\"/save/%s\" method=\"POST\">"+
	//	"<textarea name=\"body\">%s</textarea><br>"+
	//	"<input type=\"submit\" value=\"Save\">"+
	//	"</form>",
	//	p.Title, p.Title, p.Body)

	//t, _ := template.ParseFiles("../templates/edit.html")
	//t.Execute(w, p)

	render.Template(w, "edit", p)

} //EditHandler

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]

	body := r.FormValue("body")

	p := &model.Page{Title: title, Body: []byte(body)}

	p.Save()

	http.Redirect(w, r, "/view/" + title, http.StatusFound)
} //SaveHandler

