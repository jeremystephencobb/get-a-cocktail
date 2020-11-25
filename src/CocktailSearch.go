package main

import (
	"io/ioutil"
	"net/http"
)

func HandleCocktailSearch(searchType, query string) []byte {
	baseURL := handleSearchType(searchType)

	resp, err := http.Get(baseURL + query)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return bodyBytes
}
