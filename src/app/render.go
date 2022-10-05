package app

import (
	"content/src/utils"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, page *Page, err *error) {
	if err == nil || *err == nil {
		files := []string{
			"src/resources/template/" + tmpl + ".html",
			"src/resources/template/base.html",
		}
		t := template.Must(template.ParseFiles(files...))
		err := t.Execute(w, page)
		if err != nil {
			panic(err)
		}
	} else {
		utils.ThrowError(*err, w)
	}
}
