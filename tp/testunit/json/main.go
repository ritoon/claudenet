package main

import (
	"encoding/json"
	"fmt"
)

var jsonData = `{
	"Name": "Alice",
	"Age": 30000000000000000000000000000000000,
	"Email": "alice@example.com"
}`

type Person struct {
	Name  string `json:"Name"`
	Age   int    `json:"Age"` //
	Email string `json:"Email"`
}

func main() {
	// Unmarshal the JSON data into a map
	var person Person
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		panic(err)
	}

	// Print the resulting map
	fmt.Printf("%+v\n", person)
}
