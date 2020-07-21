package krollparser

import (
	"errors"
	"os"
	"path/filepath"
)

//GetFiles will look in the current working directory and
// return a slice of strings indicating all the csv files that exist there.
// If no files exist, the function will return nil, and and error
func GetFiles() ([]string, error) {
	//Create our returnable slice
	items := make([]string, 0)

	//Use the Current Working Directory
	dirname := "."
	f, err := os.Open(dirname)
	if err != nil {
		return items, err
	}

	//Get all the files
	files, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return items, err
	}

	//Loop through the files and check the extension
	for i := 0; i < len(files); i++ {
		extension := filepath.Ext(files[i])
		if files[i][0:1] != "." {
			if extension == ".csv" {
				//Add the file to our items slice
				items = append(items, files[i])
			}
		}
	}

	if len(items) < 1 {
		// return error if no files are found
		return items, errors.New("No files found")
	}

	//Return our slice of files
	return items, nil
}
