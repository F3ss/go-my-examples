package handler

import (
	"html/template"
	"net/http"
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
