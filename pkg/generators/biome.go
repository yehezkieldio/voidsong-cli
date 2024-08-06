package generators

import (
	"fmt"

	"github.com/i9ntheory/voidsong/internal/ui"
	"github.com/i9ntheory/voidsong/internal/utils"
)

type Config struct {
	Schema          string          `json:"$schema"`
	OrganizeImports OrganizeImports `json:"organizeImports"`
	Linter          Linter          `json:"linter"`
	Formatter       ConfigFormatter `json:"formatter"`
	Javascript      Javascript      `json:"javascript"`
	Vcs             Vcs             `json:"vcs"`
}

type ConfigFormatter struct {
	Enabled          bool   `json:"enabled"`
	FormatWithErrors bool   `json:"formatWithErrors"`
	LineEnding       string `json:"lineEnding"`
	IndentStyle      string `json:"indentStyle"`
	IndentWidth      int64  `json:"indentWidth"`
	LineWidth        int64  `json:"lineWidth"`
}

type Javascript struct {
	Formatter JavascriptFormatter `json:"formatter"`
}

type JavascriptFormatter struct {
	IndentStyle string `json:"indentStyle"`
	IndentWidth int64  `json:"indentWidth"`
}

type Linter struct {
	Enabled bool  `json:"enabled"`
	Rules   Rules `json:"rules"`
}

type Rules struct {
	Recommended bool        `json:"recommended"`
	Style       Style       `json:"style"`
	Correctness Correctness `json:"correctness"`
	Complexity  Complexity  `json:"complexity"`
}

type Complexity struct {
	NoStaticOnlyClass string `json:"noStaticOnlyClass"`
	NoThisInStatic    string `json:"noThisInStatic"`
}

type Correctness struct {
	NoUnusedImports string `json:"noUnusedImports"`
}

type Style struct {
	NoNonNullAssertion string `json:"noNonNullAssertion"`
}

type OrganizeImports struct {
	Enabled bool `json:"enabled"`
}

type Vcs struct {
	Enabled       bool   `json:"enabled"`
	ClientKind    string `json:"clientKind"`
	UseIgnoreFile bool   `json:"useIgnoreFile"`
	DefaultBranch string `json:"defaultBranch"`
}

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

func getConfig() (Config, error) {
	url := "https://raw.githubusercontent.com/yehezkieldio/config/master/biome.json"
	dataStruct, err := utils.GetURLContents[Config](url)
	if err != nil {
		return Config{}, err
	}

	return dataStruct, nil
}

func (g *BiomeGenerator) Run() error {
	fmt.Println(ui.InfoTextStyle.Bold(true).MarginTop(2).MarginBottom(1).Render("Configuring Biome..."))

	fmt.Println(ui.TextStyle.Render("Checking availability..."))
	avb, _ := utils.CheckAvailability()
	if avb {
		return nil
	}

	fmt.Println(ui.TextStyle.Render("Fetching configuration..."))
	config, err := getConfig()
	if err != nil {
		return err
	}

	fmt.Println(ui.TextStyle.Render("Applying configuration..."))
	err = utils.WriteJSONToFile("biome.json", config)
	if err != nil {
		return err
	}
	fmt.Println(ui.TextStyle.Render("Biome configuration applied successfully!"))

	return nil
}
