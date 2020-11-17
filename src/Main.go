package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	// "strconv"
	"strings"
)

// Generated with: https://mholt.github.io/json-to-go/
type Drinks struct {
	Drinks []struct {
		IDDrink                     string `json:"idDrink"`
		Name                        string `json:"strDrink"`
		AlternateName               string `json:"strDrinkAlternate"`
		Tags                        string `json:"strTags"`
		Video                       string `json:"strVideo"`
		Category                    string `json:"strCategory"`
		IBA                         string `json:"strIBA"`
		Alcoholic                   string `json:"strAlcoholic"`
		Glass                       string `json:"strGlass"`
		Instructions                string `json:"strInstructions"`
		DrinkThumb                  string `json:"strDrinkThumb"`
		Ingredient1                 string `json:"strIngredient1"`
		Ingredient2                 string `json:"strIngredient2"`
		Ingredient3                 string `json:"strIngredient3"`
		Ingredient4                 string `json:"strIngredient4"`
		Ingredient5                 string `json:"strIngredient5"`
		Ingredient6                 string `json:"strIngredient6"`
		Ingredient7                 string `json:"strIngredient7"`
		Ingredient8                 string `json:"strIngredient8"`
		Ingredient9                 string `json:"strIngredient9"`
		Ingredient10                string `json:"strIngredient10"`
		Ingredient11                string `json:"strIngredient11"`
		Ingredient12                string `json:"strIngredient12"`
		Ingredient13                string `json:"strIngredient13"`
		Ingredient14                string `json:"strIngredient14"`
		Ingredient15                string `json:"strIngredient15"`
		Measure1                    string `json:"strMeasure1"`
		Measure2                    string `json:"strMeasure2"`
		Measure3                    string `json:"strMeasure3"`
		Measure4                    string `json:"strMeasure4"`
		Measure5                    string `json:"strMeasure5"`
		Measure6                    string `json:"strMeasure6"`
		Measure7                    string `json:"strMeasure7"`
		Measure8                    string `json:"strMeasure8"`
		Measure9                    string `json:"strMeasure9"`
		Measure10                   string `json:"strMeasure10"`
		Measure11                   string `json:"strMeasure11"`
		Measure12                   string `json:"strMeasure12"`
		Measure13                   string `json:"strMeasure13"`
		Measure14                   string `json:"strMeasure14"`
		Measure15                   string `json:"strMeasure15"`
		StrCreativeCommonsConfirmed string `json:"strCreativeCommonsConfirmed"`
		DateModified                string `json:"dateModified"`
	} `json:"drinks"`
}

var BBlack string = "\033[1;30m"
var BRed string = "\033[1;31m"
var BGreen string = "\033[1;32m"
var BYellow string = "\033[1;33m"
var BBlue string = "\033[1;34m"
var BPurple string = "\033[1;35m"
var BCyan string = "\033[1;36m"
var BWhite string = "\033[1;37m"

func main() {
	fmt.Println(string(BRed), "Welcome to the cocktail recipe command line tool. Type `quit` at any time to exit the program")
	handleUserInput()
}

func handleUserInput() {
	var searchType string
	var query string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(string(BWhite), "Search by: `name`")
		searchType, _ = reader.ReadString('\n')
		searchType = strings.Replace(searchType, "\n", "", -1)
		if searchType != "quit" {
			prompt := fmt.Sprintf("Search for a cocktail by %v\n", searchType)
			fmt.Println(prompt)
			query, _ = reader.ReadString('\n')
			query = strings.Replace(query, "\n", "", -1)

			if query != "quit" {
				handleQuery(searchType, query)
			}
		}
		if searchType == "quit" || query == "quit" {
			fmt.Println("Goodbye")
			break

		}
	}
}

func handleQueryType(selection string) string {
	var response string
	switch {
	case strings.Contains("name", selection) || selection == "name":
		response = "https://www.thecocktaildb.com/api/json/v1/1/search.php?s="
	case strings.Contains("ing", selection) || selection == "ingredient name":
		response = "https://www.thecocktaildb.com/api/json/v1/1/search.php?i="
	case strings.Contains("ran", selection) || selection == "random":
		response = "https://www.thecocktaildb.com/api/json/v1/1/random.php"
	case strings.Contains("gla", selection) || selection == "glass type":
		response = "https://www.thecocktaildb.com/api/json/v1/1/filter.php?g="
	default:
		fmt.Println("Unrecognized query, search for cocktail by name")
		response = "https://www.thecocktaildb.com/api/json/v1/1/search.php?s="
	}
	return response
}

func handleQuery(searchType, query string) {
	var fullURL string
	baseURL := handleQueryType(searchType)
	if query == "random" {
		fullURL = baseURL
	} else {
	fullURL = baseURL + query
	}
	resp, err := http.Get(fullURL)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	buildDrinkInstructions(bodyBytes)
}

func buildDrinkInstructions(bodyBytes []byte) {
	var drinkList Drinks
	json.Unmarshal(bodyBytes, &drinkList)

	fields := reflect.TypeOf(drinkList.Drinks[0])
	values := reflect.ValueOf(drinkList.Drinks[0])

	num := fields.NumField()

	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		if value.String() != "" {
			fmt.Println(string(BPurple), field.Name, ":")
			fmt.Println(string(BCyan), value)
			fmt.Println(string(BWhite), "")
		}
	}
}
