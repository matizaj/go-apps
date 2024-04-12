package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type templateData struct {
	Data map[string]any
}

func (app *Application) render(w http.ResponseWriter, t string, td *templateData) {
	var tmpl *template.Template

	// if we are using template cache, try to get the template from map
	if app.config.useCache {
		if templateFromMap, ok := app.templateMap[t]; ok {
			tmpl = templateFromMap
		}
	}

	if tmpl == nil {
		newTemplate, err := app.buildTemplateFromDisk(t)
		if err != nil {
			log.Println("err building template ", err)
			return
		}

		log.Println("building template from disk")
		tmpl = newTemplate
	}

	if td == nil {
		td = &templateData{}
	}

	if err := tmpl.ExecuteTemplate(w, t, td); err != nil {
		log.Println("error executing template ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *Application) buildTemplateFromDisk(t string) (*template.Template, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	templateSlice := []string{
		"templates/base.layout.gohtml",
		"templates/partials/header.partial.gohtml",
		"templates/partials/footer.partial.gohtml",
		fmt.Sprintf("templates/%s", t),
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		log.Println("cant parse templates", err)
		return nil, err
	}
	app.templateMap[t] = tmpl
	return tmpl, nil
}
