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
			"v1.0.0": {
				"lastname": "Smith",
				"firstname": "Jane"
			},
			"v2.0.0": {
				"lastname": "Doe",
				"firstname": "John"
			}
		},
		"title": "The Ultimate Database Study Guide",
		"category": [
		  "Non-Fiction",
		  "Technology"
		]
	  }`

	flattnedJSON, err := flattener.FlattenJSON(jsonStr, flattener.DotSeparator, flattener.WithDepth(2))
	if err != nil {
		panic(err)
	}

	fmt.Println(flattnedJSON)
}
