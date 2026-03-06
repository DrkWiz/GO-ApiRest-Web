package services

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Chiste struct {
	Chiste string `json:"value"`
}

type SearchResponse struct {
	Total  int      `json:"total"`
	Result []Chiste `json:"result"`
}

func GetChistesRandomService() ([]byte, int, error) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	//ahora modifico el body para solo pasar el chiste al frontend
	var chiste Chiste
	if err := json.Unmarshal(body, &chiste); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	chisteJSON, err := json.Marshal(chiste)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return chisteJSON, resp.StatusCode, nil
}

func GetCategoriasService() ([]byte, int, error) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/categories")
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	//modifico el body para solo pasar las categorias al frontend
	var categorias []string
	if err := json.Unmarshal(body, &categorias); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	categoriasJSON, err := json.Marshal(categorias)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return categoriasJSON, resp.StatusCode, nil
}

func GetChistesPorCategoriasServices(category string) ([]byte, int, error) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random?category=" + url.QueryEscape(category))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	//modifico el body para solo pasar el chiste al frontend
	var chiste Chiste
	if err := json.Unmarshal(body, &chiste); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	chisteJSON, err := json.Marshal(chiste)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return chisteJSON, resp.StatusCode, nil
}

func SearchChistesPalabrasServices(keyword string) ([]byte, int, error) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/search?query=" + url.QueryEscape(keyword))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	//modifico el body para solo pasar el chiste al frontend
	var searchResponse SearchResponse
	if err := json.Unmarshal(body, &searchResponse); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	chistesJSON, err := json.Marshal(searchResponse.Result)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return chistesJSON, resp.StatusCode, nil
}
