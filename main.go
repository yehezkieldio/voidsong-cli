package main

import (
	"os"

	"github.com/i9ntheory/voidsong/cmd/voidsong"
)

func main() {
	if err := voidsong.Execute(); err != nil {
		os.Exit(1)
	}
}
