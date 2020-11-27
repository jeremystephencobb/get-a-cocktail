package main

import (
	"fmt"
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

func handleSearchType(selection string) string {
	var response string
	switch {
	case selection == "name":
		response = "https://www.thecocktaildb.com/api/json/v1/1/search.php?s="
	case selection == "random":
		response = "https://www.thecocktaildb.com/api/json/v1/1/random.php"
	case selection == "ingredient name":
		response = "https://www.thecocktaildb.com/api/json/v1/1/filter.php?i="
	case selection == "glass type":
		response = "https://www.thecocktaildb.com/api/json/v1/1/filter.php?g="
	default:
		fmt.Println("Unrecognized query, search for cocktail by name")
		response = "https://www.thecocktaildb.com/api/json/v1/1/search.php?s="
	}
	return response
}
