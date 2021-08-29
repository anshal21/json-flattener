package main

import (
	"fmt"

	flattener "github.com/anshal21/json-flattener"
)

func main() {
	jsonStr := `{
		"isbn": "123-456-222",
		"author": {
		  "lastname": "Doe",
		  "firstname": "Jane"
		},
		"editor": {
		  "lastname": "Smith",
		  "firstname": "Jane"
		},
		"title": "The Ultimate Database Study Guide",
		"category": [
		  "Non-Fiction",
		  "Technology"
		]
	  }`

	flattnedJSON, err := flattener.FlattenJSON(jsonStr, flattener.DotSeparator)
	if err != nil {
		panic(err)
	}

	fmt.Println(flattnedJSON)
}
