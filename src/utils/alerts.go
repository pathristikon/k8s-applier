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

	setContextColors(kind, text)

	os.Exit(0)
}


func setContextColors(kind string, text string) {
	var color string

	switch kind {
	case "ERR":
		color = "\u001b[31m"
	case "NOTICE":
		color = "\u001b[34m"
	case "WARNING":
		color = "\u001b[33m"
	}
	
	fmt.Printf("%s[%s] %s \u001b[0m\n", color, kind, text)
}
