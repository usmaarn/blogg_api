package utils

import (
	"fmt"
	"html/template"
	"io"
	"path"
)

func RenderTemplate(w io.Writer, templates []string, data any) {
	files := []string{}

	for _, file := range templates {
		filePath := path.Join("templates", fmt.Sprintf("%s.html", file))
		files = append(files, filePath)
	}

	tmpl, err := template.ParseFiles(files...)
	Panic(err)

	err = tmpl.Execute(w, data)
	Panic(err)
}
