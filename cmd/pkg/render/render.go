package render

import (
	"fmt"
	"html/template"

	"net/http"
)

// renderTemplate ahora recibe el nombre del archivo template como string

func RenderTemplate(w http.ResponseWriter, html string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + html)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error al ejecutar el template:", err)
		return
	}
}
