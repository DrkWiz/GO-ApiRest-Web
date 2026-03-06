package services

import (
	"io"
	"net/http"
)

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

	return body, resp.StatusCode, nil
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

	return body, resp.StatusCode, nil
}

func GetChistesPorCategoriasServices(category string) ([]byte, int, error) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random?category=" + category)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return body, resp.StatusCode, nil
}

func SearchChistesPalabrasServices(keyword string) ([]byte, int, error) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/search?query=" + keyword)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return body, resp.StatusCode, nil
}
