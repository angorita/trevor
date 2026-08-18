package main

import (
	"fmt"
	"github.com/angorita/trevor/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	fmt.Printf("Visualizá la app en: http://localhost%s/about\n", portNumber)

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
