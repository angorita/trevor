package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// renderTemplate ahora recibe el nombre del archivo template como string
func RenderTemplate(w http.ResponseWriter, html string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + html)
	if err != nil {
		fmt.Println("Error al leer el archivo template:", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error al ejecutar el template:", err)
		return
	}
}
