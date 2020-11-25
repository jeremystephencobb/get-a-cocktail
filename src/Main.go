package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// Generated with: https://mholt.github.io/json-to-go/
type Drinks struct {
	Drinks []struct {
		IDDrink      string `json:"idDrink"`
		Name         string `json:"strDrink"`
		Category     string `json:"strCategory"`
		Alcoholic    string `json:"strAlcoholic"`
		Glass        string `json:"strGlass"`
		Instructions string `json:"strInstructions"`
		Ingredient1  string `json:"strIngredient1"`
		Measure1     string `json:"strMeasure1"`
		Ingredient2  string `json:"strIngredient2"`
		Measure2     string `json:"strMeasure2"`
		Ingredient3  string `json:"strIngredient3"`
		Measure3     string `json:"strMeasure3"`
		Ingredient4  string `json:"strIngredient4"`
		Measure4     string `json:"strMeasure4"`
		Ingredient5  string `json:"strIngredient5"`
		Measure5     string `json:"strMeasure5"`
		Ingredient6  string `json:"strIngredient6"`
		Measure6     string `json:"strMeasure6"`
		Ingredient7  string `json:"strIngredient7"`
		Measure7     string `json:"strMeasure7"`
		Ingredient8  string `json:"strIngredient8"`
		Measure8     string `json:"strMeasure8"`
		Ingredient9  string `json:"strIngredient9"`
		Measure9     string `json:"strMeasure9"`
		Ingredient10 string `json:"strIngredient10"`
		Measure10    string `json:"strMeasure10"`
		Ingredient11 string `json:"strIngredient11"`
		Measure11    string `json:"strMeasure11"`
		Ingredient12 string `json:"strIngredient12"`
		Measure12    string `json:"strMeasure12"`
		Ingredient13 string `json:"strIngredient13"`
		Measure13    string `json:"strMeasure13"`
		Ingredient14 string `json:"strIngredient14"`
		Measure14    string `json:"strMeasure14"`
		Measure15    string `json:"strMeasure15"`
		Ingredient15 string `json:"strIngredient15"`
	} `json:"drinks"`
}

func main() {
	fmt.Println(string(BRed), "Welcome to the cocktail recipe command line tool. Type `quit` at any time to exit the program")
	handleUserInput()
}

func handleUserInput() {
Main:
	for {
		searchTypeprompt := promptui.Select{
			Label: "Select Day",
			Items: []string{"name",
				"ingredient name",
				"random",
				"glass type",
				"name",
				"quit",
			},
		}

		_, searchValueResult, searchValueErr := searchTypeprompt.Run()

		if searchValueErr != nil {
			fmt.Printf("Prompt failed %v\n", searchValueErr)
			return
		}

		switch {
		case searchValueResult == "quit":
			break Main
		case searchValueResult == "random":
			searchByRandom()
		case searchValueResult == "name":
			searchByName(searchValueResult)
		case searchValueResult == "ingredient name":
			searchByIngredient(searchValueResult)
		case searchValueResult == "glass type":
			searchByGlass(searchValueResult)
		default:
			fmt.Println("search by ", searchValueResult)
		}
	}
}

func searchByRandom() {
	bodyBytes := HandleCocktailSearch("random", "")
	BuildDrinkInstructions(bodyBytes)
}

func searchByName(searchValueResult string) {
	valResult, valErr := HandleTextSearch(searchValueResult)

	if valErr != nil {
		fmt.Printf("Something went wrong, please try again")
	}
	bodyBytes := HandleCocktailSearch(searchValueResult, valResult)
	BuildDrinkInstructions(bodyBytes)
}

// user selects ingredient
// user is directed to HandleTextSearch, which returns their query
// 'ingredient' and 'query' are passed to HandleCocktailSearch which returns bodyBytes
// bodyBytes is passed to BuildDrinkList, which creates a prompt to select a drink from and returns  a drink
// HandlecocktailSearch searches for the selected drink by name
// BuildDrinkInstructions builds and displays the drink instructions
func searchByIngredient(searchValueResult string) {
	valResult, valErr := HandleTextSearch(searchValueResult)

	if valErr != nil {
		fmt.Printf("Something went wrong, please try again")
	}
	bodyBytes := HandleCocktailSearch(searchValueResult, valResult)
	drinkName := BuildDrinkList(bodyBytes)
	if len(drinkName) >= 1 {
		b := HandleCocktailSearch("name", drinkName)
		BuildDrinkInstructions(b)
	} else {
		fmt.Println("There are no cocktails with that ingredient")
	}
}

func searchByGlass(searchValueResult string) {
	valResult, valErr := HandleTextSearch(searchValueResult)

	if valErr != nil {
		fmt.Printf("Something went wrong, please try again")
	}
	bodyBytes := HandleCocktailSearch(searchValueResult, valResult)
	drinkName := BuildDrinkList(bodyBytes)
	if len(drinkName) >= 1 {
		b := HandleCocktailSearch("name", drinkName)
		BuildDrinkInstructions(b)
	} else {
		fmt.Println("There are no cocktails served in that glass")
	}
}

func handleSearchType(selection string) string {
	var response string
	switch {
	case selection == "name": // ✔
		response = "https://www.thecocktaildb.com/api/json/v1/1/search.php?s="
	case selection == "random": // ✔
		response = "https://www.thecocktaildb.com/api/json/v1/1/random.php"
	case selection == "ingredient name": // ✔
		response = "https://www.thecocktaildb.com/api/json/v1/1/filter.php?i="
	case selection == "glass type":
		response = "https://www.thecocktaildb.com/api/json/v1/1/filter.php?g="
	default:
		fmt.Println("Unrecognized query, search for cocktail by name")
		response = "https://www.thecocktaildb.com/api/json/v1/1/search.php?s="
	}
	return response
}

var BBlack string = "\033[1;30m"
var BRed string = "\033[1;31m"
var BGreen string = "\033[1;32m"
var BYellow string = "\033[1;33m"
var BBlue string = "\033[1;34m"
var BPurple string = "\033[1;35m"
var BCyan string = "\033[1;36m"
var BWhite string = "\033[1;37m"
