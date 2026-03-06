package handlers

import (
	"go-apirest-web/internal/services"
	"io"
	"net/http"
)

// Home maneja la ruta raíz y muestra un mensaje de bienvenida y me informa sobre la pagina de la API de Chuck Norris
func Home(writer http.ResponseWriter, response *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	writer.WriteHeader(http.StatusOK)
	io.WriteString(writer, "<h1>Bienvenido a la API de Chuck Norris</h1>")
}

// getChistesRandom maneja la ruta /random y muestra un chiste aleatorio de Chuck Norris y lo manda al frontend
func GetChistesRandom(writer http.ResponseWriter, response *http.Request) {

	body, status, err := services.GetChistesRandomService()
	if err != nil {
		http.Error(writer, "Error al obtener chiste aleatorio", status)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(body)
}

// getCategoria maneja la ruta /categoria y muestra las categorias disponibles de chistes de Chuck Norris y lo manda al frontend
func GetCategoria(writer http.ResponseWriter, response *http.Request) {
	body, status, err := services.GetCategoriasService()
	if err != nil {
		http.Error(writer, "Error al obtener categorias", status)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(body)
}

// GetChistesCategoria maneja la ruta /categoria/{categoria} y muestra un chiste de Chuck Norris de la categoria seleccionada y lo manda al frontend
func GetChistesCategoria(writer http.ResponseWriter, response *http.Request) {
	category := response.URL.Query().Get("category")
	if category == "" {
		http.Error(writer, "parámetro 'category' es obligatorio", http.StatusBadRequest)
		return
	}

	body, status, err := services.GetChistesPorCategoriasServices(category)
	if err != nil {
		http.Error(writer, "Error al obtener chiste por categoría", status)
		return
	}

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(status)
	if _, err := writer.Write(body); err != nil {
		http.Error(writer, "Error al escribir respuesta", http.StatusInternalServerError)
		return
	}
}

// GetChistesPalabra maneja la ruta /search?query={palabra} y muestra un chiste de Chuck Norris que contenga la palabra clave seleccionada y lo manda al frontend
func GetChistesPalabra(writer http.ResponseWriter, response *http.Request) {
	keyword := response.URL.Query().Get("query")
	if keyword == "" {
		http.Error(writer, "parámetro 'query' es obligatorio", http.StatusBadRequest)
		return
	}

	body, status, err := services.SearchChistesPalabrasServices(keyword)
	if err != nil {
		http.Error(writer, "Error al buscar chistes por palabra clave", status)
		return
	}

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(status)
	if _, err := writer.Write(body); err != nil {
		http.Error(writer, "Error al escribir respuesta", http.StatusInternalServerError)
		return
	}

}
