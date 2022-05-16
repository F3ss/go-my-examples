package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, templateName string) error {
	parseTemplate, err := template.ParseFiles("./templates/" + templateName)
	if err != nil {
		return err
	}

	if err := parseTemplate.Execute(w, nil); err != nil {
		return err
	}

	return nil
}

func RenderTemplateTest(w http.ResponseWriter, tmpl string) error {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return err
		}
	}

	return nil
}
