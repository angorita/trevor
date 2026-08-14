package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

// Home es el manejador para la página de inicio
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

// About es el manejador para la página sobre nosotros
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

// renderTemplate ahora recibe el nombre del archivo template como string
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
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

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	fmt.Printf("Visualizá la app en: http://localhost%s/about\n", portNumber)

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
