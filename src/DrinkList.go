package main

import (
	"encoding/json"
	"fmt"
	"github.com/manifoldco/promptui"
	"strings"
)

type DrinkList struct {
	Drinks []struct {
		StrDrink      string `json:"strDrink"`
		StrDrinkThumb string `json:"strDrinkThumb"`
		IDDrink       string `json:"idDrink"`
	} `json:"drinks"`
}

func BuildDrinkList(bodyBytes []byte) string {
	var drinkList DrinkList
	json.Unmarshal(bodyBytes, &drinkList)

	n := len(drinkList.Drinks)

	// display the list of drinks
	var drinkchoices []string

	for i := 0; i < n; i++ {
		drinkchoices = append(drinkchoices, drinkList.Drinks[i].StrDrink)
	}

	if len(drinkchoices) >= 1 {
		searchTypeprompt := promptui.Select{
			Label: "Select Drink",
			Items: drinkchoices,
		}

		_, result, err := searchTypeprompt.Run()

		if err != nil {
			fmt.Println(err)
		}

		return strings.ReplaceAll(result, " ", "_")
	}

	return ""
}
