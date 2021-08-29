
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/anshal21/json-flattener/blob/main/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/anshal21/json-flattener)](https://goreportcard.com/report/github.com/anshal21/json-flattener)

# json-flattener
json-flattener helps in flattening complex nested JSONs in different ways
 
 ## Installation
 ```bash
 go get github.com/anshal21/json-flattener
 ```

 ## Examples

 ### 1. [ Simple Flattening ](https://github.com/anshal21/json-flattener/tree/main/examples/simple_flatten)  
 ```go
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
 ``` 
```json
    {
        "author.firstname": "Jane",
        "author.lastname": "Doe",
        "category.0": "Non-Fiction",
        "category.1": "Technology",
        "editor.firstname": "Jane",
        "editor.lastname": "Smith",
        "isbn": "123-456-222",
        "title": "The Ultimate Database Study Guide"
    }
```  

 ### 2. [ Flatten and ignore arrays ](https://github.com/anshal21/json-flattener/tree/main/examples/flatten_ignore_array)  
 ```go
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

	flattnedJSON, err := flattener.FlattenJSON(jsonStr, flattener.DotSeparator, flattener.IgnoreArray())
 ``` 
```json
    {
        "author.firstname": "Jane",
        "author.lastname": "Doe",
        "category": [
            "Non-Fiction",
            "Technology"
        ],
        "editor.firstname": "Jane",
        "editor.lastname": "Smith",
        "isbn": "123-456-222",
        "title": "The Ultimate Database Study Guide"
    }
```  


 ### 3. [ Flatten till specified depth ](https://github.com/anshal21/json-flattener/tree/main/examples/simple_flatten)  
 ```go
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
 ``` 
```json
    {
        "author.firstname": "Jane",
        "author.lastname": "Doe",
        "category.0": "Non-Fiction",
        "category.1": "Technology",
        "editor.v1.0.0": {
            "firstname": "Jane",
            "lastname": "Smith"
        },
        "editor.v2.0.0": {
            "firstname": "John",
            "lastname": "Doe"
        },
        "isbn": "123-456-222",
        "title": "The Ultimate Database Study Guide"
    }
```  