package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

func main() {
	myJson := `
[
	{
		"first_name": "Clark",
		"second_name": "Clarkovskii"
	},
	{
		"first_name": "Tod",
		"second_name": "Tododododo"
	}
]`
	persons, err := unmarshalPersons(myJson)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(marshalPersons(&persons))
}

func marshalPersons(persons *[]Person) (string, error) {
	var result, err = json.MarshalIndent(persons, "", "	   ")

	if err != nil {
		return "", err
	}

	return string(result), nil
}

func unmarshalPersons(persons string) ([]Person, error) {
	var result []Person

	if err := json.Unmarshal([]byte(persons), &result); err != nil {
		return nil, err
	}

	return result, nil
}
