package main

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"strings"
)

func HandleTextSearch(searchValueResult string) (string, error) {
	searchValueprompt := promptui.Prompt{
		Label: "Search:",
	}

	fmt.Printf("Search by: %q\n", searchValueResult)
	result, err := searchValueprompt.Run()

	if err != nil {
		return "", errors.New("Something went wrong, please try again")
	}
	return strings.ReplaceAll(result, " ", "_"), nil
}
