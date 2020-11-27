package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

// Generated with: https://mholt.github.io/json-to-go/
type Drinks struct {
	Drinks []struct {
		Name         string `json:"strDrink"`
		Glass        string `json:"strGlass"`
		Measure1     string `json:"strMeasure1"`
		Ingredient1  string `json:"strIngredient1"`
		Measure2     string `json:"strMeasure2"`
		Ingredient2  string `json:"strIngredient2"`
		Measure3     string `json:"strMeasure3"`
		Ingredient3  string `json:"strIngredient3"`
		Measure4     string `json:"strMeasure4"`
		Ingredient4  string `json:"strIngredient4"`
		Measure5     string `json:"strMeasure5"`
		Ingredient5  string `json:"strIngredient5"`
		Measure6     string `json:"strMeasure6"`
		Ingredient6  string `json:"strIngredient6"`
		Measure7     string `json:"strMeasure7"`
		Ingredient7  string `json:"strIngredient7"`
		Measure8     string `json:"strMeasure8"`
		Ingredient8  string `json:"strIngredient8"`
		Measure9     string `json:"strMeasure9"`
		Ingredient9  string `json:"strIngredient9"`
		Measure10    string `json:"strMeasure10"`
		Ingredient10 string `json:"strIngredient10"`
		Measure11    string `json:"strMeasure11"`
		Ingredient11 string `json:"strIngredient11"`
		Measure12    string `json:"strMeasure12"`
		Ingredient12 string `json:"strIngredient12"`
		Measure13    string `json:"strMeasure13"`
		Ingredient13 string `json:"strIngredient13"`
		Measure14    string `json:"strMeasure14"`
		Ingredient14 string `json:"strIngredient14"`
		Measure15    string `json:"strMeasure15"`
		Ingredient15 string `json:"strIngredient15"`
		Instructions string `json:"strInstructions"`
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
				"combo search",
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
		case searchValueResult == "combo search":
			comboSearch()
		default:
			fmt.Println("search by ", searchValueResult)
		}
	}
}

func comboSearch() {
	ingredientDrinkChoices := comboIngredient()
	glassDrinkChoices := comboGlass()
	i := len(ingredientDrinkChoices)
	g := len(glassDrinkChoices)

	var comboCocktails []string
	for a := 0; a < i; a++ {
		for b := 0; b < g; b++ {
			if ingredientDrinkChoices[a] == glassDrinkChoices[b] {
				comboCocktails = append(comboCocktails, glassDrinkChoices[b])
			}
		}
	}

	if len(comboCocktails) >= 1 {
		// create a prompt with the cocktail names
		searchTypeprompt := promptui.Select{
			Label: "Select Drink",
			Items: comboCocktails,
		}
	
		_, result, err := searchTypeprompt.Run()
	
		if err != nil {
			panic(err)
		}
		b := HandleCocktailSearch("name", strings.ReplaceAll(result, " ", "_"))
		BuildDrinkInstructions(b)
	} else {
		fmt.Println("No drinks for that combo")
	}
}

func comboGlass() []string {
	// pick a glass
	valResult, valErr := HandleTextSearch("glass type:")

	if valErr != nil {
		fmt.Printf("Something went wrong, please try again")
	}
	gbodyBytes := HandleCocktailSearch("glass type", valResult)
	var glassDrinkList DrinkList
	json.Unmarshal(gbodyBytes, &glassDrinkList)

	l := len(glassDrinkList.Drinks)

	var glassDrinkchoices []string

	for i := 0; i < l; i++ {
		glassDrinkchoices = append(glassDrinkchoices, glassDrinkList.Drinks[i].Name)
	}
	// build up a list of cocktails available
	return glassDrinkchoices
}

func comboIngredient() []string {
	valResult, valErr := HandleTextSearch("ingredient name:")

	if valErr != nil {
		fmt.Printf("Something went wrong, please try again")
	}
	bodyBytes := HandleCocktailSearch("ingredient name", valResult)
	var drinkList DrinkList
	json.Unmarshal(bodyBytes, &drinkList)

	n := len(drinkList.Drinks)

	var drinkchoices []string

	for i := 0; i < n; i++ {
		drinkchoices = append(drinkchoices, drinkList.Drinks[i].Name)
	}

	return drinkchoices
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

var BBlack string = "\033[1;30m"
var BRed string = "\033[1;31m"
var BGreen string = "\033[1;32m"
var BYellow string = "\033[1;33m"
var BBlue string = "\033[1;34m"
var BPurple string = "\033[1;35m"
var BCyan string = "\033[1;36m"
var BWhite string = "\033[1;37m"
