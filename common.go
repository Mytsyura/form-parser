package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// FillForm asks the user to fill in the form by answering each question (handling conditional questions)
func FillForm(form *Form) map[string]string {
	results := make(map[string]string)

	fmt.Printf("Filling form: %s\n", form.Title)

	for _, q := range form.Questions {
		askQuestion(q, results)
	}

	return results
}

// askQuestion handles asking the question and any follow-up conditional questions
func askQuestion(q Question, results map[string]string) {
	fmt.Printf("Question: %s (%s) ", q.Label, q.Type)
	if len(q.Options) > 0 {
		fmt.Printf("Options: %v ", q.Options)
	}
	if q.DefaultValue != "" {
		fmt.Printf("[default: %s]: ", q.DefaultValue)
	}

	var answer string
	fmt.Scanln(&answer)

	if answer == "" && q.DefaultValue != "" {
		answer = q.DefaultValue
	}

	results[q.Label] = answer

	// If there are follow-up questions, ask them based on the user's answer
	if followUpQuestions, exists := q.Next[answer]; exists {
		for _, nextQuestion := range followUpQuestions {
			askQuestion(nextQuestion, results)
		}
	}
}

// LoadFormFromFile loads a form from a given file path (JSON or YAML)
func LoadFormFromFile(filename string) (*Form, error) {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var form Form
	if err := json.Unmarshal(fileData, &form); err == nil {
		return &form, nil
	}

	if err := yaml.Unmarshal(fileData, &form); err == nil {
		return &form, nil
	}

	return nil, errors.New("file format not supported or incorrect structure")
}
