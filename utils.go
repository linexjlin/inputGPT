package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"
)

func showAsJson(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("Error converting data to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func applyStringTemplate(temp string, vars map[string]string) string {
	tmpl, err := template.New("tmpl").Parse(temp)
	if err != nil {
		fmt.Println(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, vars)
	if err != nil {
		fmt.Println(err)
	}

	result := buf.String()
	return result
}
