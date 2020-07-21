package main

import (
	"fmt"
	"kroll-data-stripper/krollparser"
	"log"
	"os"
)

func main() {
	//Start by getting a list of all files in the current working directory ending with .csv
	items, err := krollparser.GetFiles()
	if err != nil {
		log.Fatal(err)
	}

	var selectedFile string = ""

	if len(items) == 1 {
		//If single file exists, confirm that the user wishes to process the file
		if err = krollparser.ConfirmFile(items[0]); err != nil {
			log.Fatal(err)
		}
		selectedFile = items[0]
	} else {
		//If multiple files exist, prompt the user to see which one they'd like to use
		selectedFile, err = krollparser.SelectFile(items)
		if err != nil {
			log.Fatal(err)
		}
	}

	//Output selected file to user
	fmt.Println("Now Processing File: ", selectedFile)
	fmt.Println("")
	fmt.Println("")
	// Now that we have a file selected, load the file into a slice of structs
	records, err := krollparser.ReadKrollCSV(selectedFile)
	if err != nil {
		log.Fatal(err)
	}

	// Let's write the slice of structs to a CSV output file
	err = krollparser.MakeCSV(records)
	if err != nil {
		log.Fatal(err)
	}

	// Let's prefix the selected file with "delete-" so the user knows it's safe to delete the file now
	newFileName := fmt.Sprintf("delete-%s", selectedFile)
	fmt.Println("Renaming ", selectedFile, " to ", newFileName)
	err = os.Rename(selectedFile, newFileName)
	if err != nil {
		log.Fatal(err)
	}

	krollparser.EndRuntime()
}
