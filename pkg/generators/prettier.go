package generators

import (
	"fmt"

	"github.com/i9ntheory/voidsong/internal/ui"
	"github.com/i9ntheory/voidsong/internal/utils"
)

type PrettierGenerator struct{}

func (g *PrettierGenerator) Name() string {
	return "Prettier"
}

func (g *PrettierGenerator) Description() string {
	return "Generate my opinionated Prettier configuration."
}

func (g *PrettierGenerator) FilterValue() string {
	return "prettier"
}

func (g *PrettierGenerator) Run() error {
	fmt.Println(ui.InfoTextStyle.Bold(true).MarginTop(2).MarginBottom(1).Render("Configuring Prettier..."))

	fmt.Println(ui.TextStyle.Render("Checking availability..."))
	avb, _ := utils.CheckAvailability()
	if avb {
		return nil
	}

	configFilesRegex := `\.prettierrc|\.prettierrc\.js|\.prettierrc\.json|\.prettierrc\.yaml|\.prettierrc\.yml|prettier\.config\.js|prettier\.config\.json|prettier\.config\.yaml|prettier\.config\.yml`
	if utils.CheckConfigurationFile(configFilesRegex) {
		fmt.Println(ui.ErrorTextStyle.Render("Please remove any existing Prettier configuration files before running this command."))
		return nil
	}

	return nil
}
