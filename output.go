package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// PrintResults prints the form results to the terminal
func PrintResults(results map[string]string) {
	fmt.Println("Form Results:")
	for question, answer := range results {
		fmt.Printf("%s: %s\n", question, answer)
	}
}

// SaveResultsToFile saves the results to a specified file in JSON format
func SaveResultsToFile(filename string, results map[string]string) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}
