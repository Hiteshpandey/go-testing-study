package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "home.page.gohtml", &TemplateData{})
	fmt.Fprintln(w, "This is Home page")
}

type TemplateData struct {
	Ip   string
	Data map[string]any
}

var templatesPath = "./templates/"

func (app *application) render(w http.ResponseWriter, r *http.Request, tplt string, data *TemplateData) error {
	// parse template from disk
	parsedTemplate, err := template.ParseFiles(path.Join(templatesPath, tplt))
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return err
	}

	data.Ip = app.ipFromContext(r.Context())

	err = parsedTemplate.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}
