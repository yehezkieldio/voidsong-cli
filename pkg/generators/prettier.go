package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/i9ntheory/voidsong/internal/ui"
)

type PrettierGenerator struct{}

func (g *PrettierGenerator) Name() string {
	return "Prettier"
}

func (g *PrettierGenerator) FilterValue() string {
	return "prettier"
}

func (g *PrettierGenerator) Run() error {
	fmt.Println("Finding package.json...")
	if _, err := os.Stat("package.json"); os.IsNotExist(err) {
		fmt.Println(" Cannot find package.json!")
		fmt.Println(ui.ErrorTextStyle.Render("\nPlease ensure you are in the root of a project with a package.json file."))
		return nil
	} else {
		fmt.Println(" Found package.json!")
	}

	fmt.Println("Finding existing configuration...")
	matches, err := filepath.Glob("*prettierrc*")
	if err != nil {
		fmt.Println(ui.ErrorTextStyle.Render("\nError finding existing configuration: " + err.Error()))
		return nil
	}

	if len(matches) > 0 {
		fmt.Println(" Existing configuration found!")
		fmt.Println(" Please remove the existing configuration before running this generator.")
		fmt.Println("  - " + strings.Join(matches, "\n  - "))
	} else {
		fmt.Println(" No existing configuration found, creating new configuration...")
	}

	return nil
}
