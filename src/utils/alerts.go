package utils

import (
	"fmt"
	"os"
)

func Alert(kind string, text string) {
	choices := map[string]bool{"ERR": true, "NOTICE": true, "WARNING": true}

	if _, validChoice := choices[kind]; !validChoice {
		panic("[DEBUG] You requested an invalid alert type! \n")
	}

	fmt.Printf("[%s] %s \n", kind, text)

	os.Exit(0)
}
