package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {

	var dummy string
	dummy = `{
        "name": "Tolexo Online Pvt. Ltd",
        "age_in_years": 8.5,
        "origin": "Noida",
        "head_office": "Noida, Uttar Pradesh",
        "address": [
            {
                "street": "91 Springboard",
                "landmark": "Axis Bank",
                "city": "Noida",
                "pincode": 201301,
                "state": "Uttar Pradesh"
            },
            {
                "street": "91 Springboard",
                "landmark": "Axis Bank",
                "city": "Noida",
                "pincode": 201301,
                "state": "Uttar Pradesh"
            }
        ],
        "sponsors": {
            "name": "One"
        },
        "revenue": "19.8 million$",
        "no_of_employee": 630,
        "str_text": ["one", "two"],
        "int_text": [1, 3, 4]
    }`

	// Unmarshal the JSON data into a map

	mp := make(map[string]interface{})
	err := json.Unmarshal([]byte(dummy), &mp)

	if err != nil {
		return
	} else {

		display(mp)
	}
}

func display(mp map[string]interface{}) {

	for i, val := range mp {
		// Getting reflection value of the current value
		ref := reflect.ValueOf(val)

		switch ref.Kind() {
		case reflect.String:
			fmt.Printf("key is :%v , value is %v , type is %v\n", i, ref, ref.Kind())

		case reflect.Map:
			// If the value is a map, recursively call print to handle nested maps
			nestedmap := ref.Interface().(map[string]interface{})
			print(nestedmap)

		case reflect.Slice:
			// If the value is a slice, iterate over its elements
			nestedslice := ref.Interface().([]interface{})
			for i, val := range nestedslice {
				value := reflect.ValueOf(val)
				if value.Kind() == reflect.Map {
					nestedMap := value.Interface().(map[string]interface{})
					print(nestedMap)
				} else {
					// If the element is not a map, print its key, value, and type
					fmt.Printf("key is :%v, value is %v, type is %v\n", i, value, value.Kind())
				}
			}

		default:
			fmt.Printf("key is :%v , value is %v , type is %v\n", i, ref, ref.Kind())
		}
	}
}
