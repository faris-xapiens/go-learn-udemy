package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Ade",
			"last_name": "Sunhaki",
			"hair_color": "Grey",
			"has_dog": true
		},
		{
			"first_name": "Kamil",
			"last_name": "Sunhaki",
			"hair_color": "Red",
			"has_dog": false
		}
	]`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Println("unmarshalled: %v", unmarshalled)

	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "Red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Michael"
	m2.LastName = "East"
	m2.HairColor = "Purple"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")

	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))

}
