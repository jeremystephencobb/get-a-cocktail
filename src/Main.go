package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

type DrinkList struct {
	Drinks []struct {
		StrDrink      string `json:"strDrink"`
		StrDrinkThumb string `json:"strDrinkThumb"`
		IDDrink       string `json:"idDrink"`
	} `json:"drinks"`
}

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
		Measure1                    string `json:"strMeasure1"`
		Ingredient2                 string `json:"strIngredient2"`
		Measure2                    string `json:"strMeasure2"`
		Ingredient3                 string `json:"strIngredient3"`
		Measure3                    string `json:"strMeasure3"`
		Ingredient4                 string `json:"strIngredient4"`
		Measure4                    string `json:"strMeasure4"`
		Ingredient5                 string `json:"strIngredient5"`
		Measure5                    string `json:"strMeasure5"`
		Ingredient6                 string `json:"strIngredient6"`
		Measure6                    string `json:"strMeasure6"`
		Ingredient7                 string `json:"strIngredient7"`
		Measure7                    string `json:"strMeasure7"`
		Ingredient8                 string `json:"strIngredient8"`
		Measure8                    string `json:"strMeasure8"`
		Ingredient9                 string `json:"strIngredient9"`
		Measure9                    string `json:"strMeasure9"`
		Ingredient10                string `json:"strIngredient10"`
		Measure10                   string `json:"strMeasure10"`
		Ingredient11                string `json:"strIngredient11"`
		Measure11                   string `json:"strMeasure11"`
		Ingredient12                string `json:"strIngredient12"`
		Measure12                   string `json:"strMeasure12"`
		Ingredient13                string `json:"strIngredient13"`
		Measure13                   string `json:"strMeasure13"`
		Ingredient14                string `json:"strIngredient14"`
		Measure14                   string `json:"strMeasure14"`
		Measure15                   string `json:"strMeasure15"`
		Ingredient15                string `json:"strIngredient15"`
		StrCreativeCommonsConfirmed string `json:"strCreativeCommonsConfirmed"`
		DateModified                string `json:"dateModified"`
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
			handleCocktailSearch(searchValueResult, "")
		case searchValueResult == "name":
			valResult, valErr := handleTextSearch(searchValueResult)

			if valErr != nil {
				fmt.Printf("Something went wrong, please try again")
			}
			bodyBytes := handleCocktailSearch(searchValueResult, valResult)
			buildDrinkInstructions(bodyBytes)

		case searchValueResult == "ingredient name":
			valResult, valErr := handleTextSearch(searchValueResult)

			if valErr != nil {
				fmt.Printf("Something went wrong, please try again")
			}
			bodyBytes := handleCocktailSearch(searchValueResult, valResult)
			drinkName := buildDrinkList(bodyBytes)

			b := handleCocktailSearch("name", strings.ReplaceAll(drinkName, " ", "_"))
			buildDrinkInstructions(b)
		default:
			fmt.Println("search by ", searchValueResult)
		}
	}
}

func handleTextSearch(searchValueResult string) (string, error) {
	searchValueprompt := promptui.Prompt{
		Label: "Search:",
	}

	fmt.Printf("Search by: %q\n", searchValueResult)
	result, err := searchValueprompt.Run()

	if err != nil {
		return "", errors.New("Something went wrong, please try again")
	}
	return result, nil
}

func handleCocktailSearch(searchType, query string) []byte {
	baseURL := handleSearchType(searchType)

	resp, err := http.Get(baseURL + query)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return bodyBytes
}

func handleIngredientSearch(selection string) {

}

func handleSearchType(selection string) string {
	var response string
	switch {
	case selection == "name": // ✔
		response = "https://www.thecocktaildb.com/api/json/v1/1/search.php?s="
	case selection == "random": // ✔
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

func buildDrinkList(bodyBytes []byte) string {
	var drinkList DrinkList
	json.Unmarshal(bodyBytes, &drinkList)

	n := len(drinkList.Drinks)

	// display the list of drinks
	var drinkchoices []string

	for i := 0; i < n; i++ {
		drinkchoices = append(drinkchoices, drinkList.Drinks[i].StrDrink)
	}

	searchTypeprompt := promptui.Select{
		Label: "Select Drink",
		Items: drinkchoices,
	}

	_, result, err := searchTypeprompt.Run()

	if err != nil {
		fmt.Println(err)
	}

	return result
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

var BBlack string = "\033[1;30m"
var BRed string = "\033[1;31m"
var BGreen string = "\033[1;32m"
var BYellow string = "\033[1;33m"
var BBlue string = "\033[1;34m"
var BPurple string = "\033[1;35m"
var BCyan string = "\033[1;36m"
var BWhite string = "\033[1;37m"
