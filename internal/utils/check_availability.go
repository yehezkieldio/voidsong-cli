package utils

import (
	"fmt"

	"github.com/i9ntheory/voidsong/internal/ui"
)

func CheckAvailability() (bool, error) {
	l := CheckPackageJSON()
	if !l {
		fmt.Println(ui.ErrorTextStyle.Render("- Can't find a package.json file in the current directory! Are you sure you're in the right place?"))
		return true, nil
	}
	fmt.Println(ui.TextStyle.Render("- Found package.json file!"))

	b := CheckBunProject()
	if !b {
		fmt.Println(ui.ErrorTextStyle.Render("- Currently I only support projects that use Bun! Please make sure you have Bun installed and using it in your project."))
		return true, nil
	}
	fmt.Println(ui.TextStyle.Render("- Found bun.lockb file!"))

	return false, nil
}
