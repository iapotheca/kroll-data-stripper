package krollparser

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

//ConfirmFile is passed a filename and returns the
func ConfirmFile(file string) error {
	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("One file found, Would you like to process this file?  %s", file),
		IsConfirm: true,
	}

	_, err := prompt.Run()

	if err != nil {
		return err
	}

	return nil
}
