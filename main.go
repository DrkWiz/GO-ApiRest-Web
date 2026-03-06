// obtención de datos de la API de Chuck Norris y renderizado de html
package main

import (
	"log"
	"net/http"

	"go-apirest-web/internal/handlers"
)

func main() {
	//dedinir rutas
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/api/random", handlers.GetChistesRandom)
	http.HandleFunc("/api/categoria", handlers.GetCategoria)
	http.HandleFunc("/api/categoria/", handlers.GetChistesCategoria)
	http.HandleFunc("/api/search", handlers.GetChistesPalabra)

	//iniciar servidor
	log.Println("Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error iniciando servidor: %v", err)
	}
}
