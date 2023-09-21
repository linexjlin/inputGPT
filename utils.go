package main

import (
	"encoding/json"
	"fmt"
)

func showAsJson(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("Error converting data to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}
