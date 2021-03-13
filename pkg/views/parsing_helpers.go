package views

import (
	"embed"
	"text/template"
)

// TODO: look more into the performance as I'm not sure about this/
// it seems cool so I'm trying it out

//go:embed layout/* pages/*
var htmlTemplates embed.FS

func parse(page string) *template.Template {
	return mustWithError(
		template.New("layout.html").ParseFS(htmlTemplates, "layout/layout.html", page))
}

//MustWithError
func mustWithError(t *template.Template, err error) *template.Template {
	if err != nil {
		errorPage, _ := template.New("layout.html").ParseFS(htmlTemplates, "layout/layout.html", "pages/error.html")

		return errorPage
	}
	return t
}
