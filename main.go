package main

import (
	"fmt"
	"os"
)

func main() {
	var importedForms []*Form
	var choice int

	fmt.Println("Welcome!")
	for {
		// Ask the user whether to import a form or use a predefined one
		fmt.Println("choose an action")
		fmt.Println("1. import a form")
		fmt.Println("2. fill in a form")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var filePath string
			fmt.Print("Enter the path to the form file (JSON or YAML): ")
			fmt.Scanln(&filePath)

			form, err := LoadFormFromFile(filePath)
			if err != nil {
				fmt.Printf("Error loading form: %v\n", err)
				continue // Go back to main menu
			}
			importedForms = append(importedForms, form)
			fmt.Printf("Form '%s' successfully imported!\n", form.Title)
		case 2:
			if len(importedForms) == 0 {
				fmt.Println("No forms imported yet.")
				continue
			}
			worker(importedForms)
			os.Exit(0)

		default:
			fmt.Println("Invalid choice. Please select a valid option: 1 or 2.")
		}
	}
}

func worker(importedForms []*Form) {
	// Only allow user to select from the imported forms, not predefined forms
	form := selectFormFromImported(importedForms)

	// Fill the selected form
	results := FillForm(form)

	// Ask user whether to print or save results
	var outputOption string
	fmt.Print("Do you want to print the results or save to a file? (print/save): ")
	fmt.Scanln(&outputOption)
	for {
		if outputOption == "print" {
			PrintResults(results)
			break
		} else if outputOption == "save" {
			var filename string
			fmt.Print("Enter the output filename: ")
			fmt.Scanln(&filename)
			if err := SaveResultsToFile(filename, results); err != nil {
				fmt.Printf("Error saving results: %v\n", err)
			}
			break
		} else {
			fmt.Println("Invalid option. Select 'print' or 'save'.")
		}
	}

}

func selectFormFromImported(importedForms []*Form) *Form {
	for {
		fmt.Println("\nSelect an imported form to fill:")
		for i, form := range importedForms {
			fmt.Printf("%d. %s\n", i+1, form.Title)
		}

		var choice int
		fmt.Scanln(&choice)

		if choice >= 1 && choice <= len(importedForms) {
			return importedForms[choice-1] // Return the selected form
		}

		fmt.Println("Invalid selection. Please select a valid form number.")
	}
}
