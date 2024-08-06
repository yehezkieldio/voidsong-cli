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

func (g *PrettierGenerator) FilterValue() string {
	return "prettier"
}

func (g *PrettierGenerator) Run() error {
	fmt.Println(ui.InfoTextStyle.Bold(true).Render("Configuring Prettier..."))

	l := utils.CheckPackageJSON()
	if !l {
		fmt.Println(ui.ErrorTextStyle.Render("Can't find a package.json file in the current directory! Are you sure you're in the right place?"))
		return nil
	}

	return nil
}
