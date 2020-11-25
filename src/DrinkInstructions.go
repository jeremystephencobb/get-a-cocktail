package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func BuildDrinkInstructions(bodyBytes []byte) {
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

// arr1 := []string{
// 	"one", "two", "three", "four", "six",
// }

// arr2 := []string{
// 	"one", "two", "four",
// }

// var commonelements []string
// for i := 0; i < len(arr1); i++ {
// 	for j := 0; j < len(arr2); j++ {
// 		if arr1[i] == arr2[j] {
// 			commonelements = append(commonelements, arr1[i])
// 		}
// 	}
// }
