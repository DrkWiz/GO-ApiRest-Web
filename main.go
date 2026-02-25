// obtencion de datos de la api de mercado libre y renderizado de html
package main

import (
	"log"
	"net/http"

	"github.com/DrkWiz/GO-ApiRest-Web/internal/handlers"
)

func main() {
	//dedinir rutas
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/api/productos", handlers.GetProductos)

	//iniciar servidor
	log.Println("Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error iniciando servidor: %v", err)
	}
}
