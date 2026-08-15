package main

import (
	"net/http"
)

// Home es el manejador para la página de inicio
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

// About es el manejador para la página sobre nosotros
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}
