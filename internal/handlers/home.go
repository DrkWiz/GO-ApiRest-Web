package handlers

import (
	"fmt"
	"net/http"
)

func Home(writer http.ResponseWriter, response *http.Request) {
	fmt.Fprintln(writer, "Wlecome to the APIREST of MeLi")
}

func GetProductos(writer http.ResponseWriter, response *http.Request) {
	fmt.Fprintln(writer, "GetProducts")
}
