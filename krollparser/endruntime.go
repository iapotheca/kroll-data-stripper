package krollparser

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

//EndRuntime just holds the window open until the user hits enter,
// that way they can see the results rather that it flashing away to quick to see
func EndRuntime() {
	prompt := promptui.Prompt{
		Label:     fmt.Sprint("Process complete.  Close window"),
		IsConfirm: true,
	}

	prompt.Run()

	return
}
