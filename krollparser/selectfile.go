package krollparser

import "github.com/manifoldco/promptui"

//SelectFile will take a slice of files,
// It will prompt the user to select one
// If the user makes a selection, it will return the file name and nil
// If the user ends execution, it return a zero-length string and an error
func SelectFile(files []string) (string, error) {

	prompt := promptui.Select{
		Label: "Please choose the file you would like to process file",
		Items: files,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}
