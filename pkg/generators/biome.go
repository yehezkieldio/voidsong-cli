package generators

import (
	"fmt"

	"github.com/i9ntheory/voidsong/internal/ui"
	"github.com/i9ntheory/voidsong/internal/utils"
)

type BiomeGenerator struct{}

func (g *BiomeGenerator) Name() string {
	return "Biome"
}

func (g *BiomeGenerator) Description() string {
	return "Generate my opinionated Biome configuration."
}

func (g *BiomeGenerator) FilterValue() string {
	return "Biome"
}

func (g *BiomeGenerator) Run() error {
	fmt.Println(ui.InfoTextStyle.Bold(true).MarginTop(2).MarginBottom(1).Render("Configuring Biome..."))

	fmt.Println(ui.TextStyle.Render("Checking availability..."))
	avb, _ := utils.CheckAvailability()
	if avb {
		return nil
	}

	return nil
}
