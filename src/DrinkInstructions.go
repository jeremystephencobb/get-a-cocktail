package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

func BuildDrinkInstructions(bodyBytes []byte) {
	var drinkList Drinks
	json.Unmarshal(bodyBytes, &drinkList)

	fields := reflect.TypeOf(drinkList.Drinks[0])
	values := reflect.ValueOf(drinkList.Drinks[0])

	num := fields.NumField()
	ingredientsTable := table.NewWriter()
	
	ingredientsTable.SetOutputMirror(os.Stdout)
	
	ingredientsTable.AppendHeader(table.Row{"Ingredient", "Amount"})
	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		if value.String() != "" {
			if field.Name != "Instructions" {
				ingredientsTable.AppendRow([]interface{}{field.Name, value})
			} 
		}
	}
	
	ingredientsTable.Render()
	fmt.Println("\n")
	fmt.Println(string(BCyan))
	fmt.Printf("Instructions: %s", values.Field(num -1))
	fmt.Println(string(BWhite))
	
}