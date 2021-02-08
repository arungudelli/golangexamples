package main

import (
	"fmt"
)

func main() {

	//creating a map
	goMap := map[string]int{"key1": 100, "key2": 200, "key3": 300}

	//Indexing the map with the key.
	value, isMapContainsKey := goMap["key2"]
	//isMapContainsKey will be true if the key contains in goMap
	if isMapContainsKey {
		//key exist
		fmt.Println("Map contains the key and the value is =  ", value)
	} else {
		//key does not exist
		fmt.Println("Map does not contains the key")
	}
}
