package handlers

import (
	"github.com/angorita/trevor/pkg/render"
	"net/http"
)

// Home es el manejador para la página de inicio
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

// About es el manejador para la página sobre nosotros
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}
