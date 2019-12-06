package utils

import (
	"fmt"
	"os"
)

func PrintHelp() {
	message := `Kubernetes cluster helper

Usage:
	` + ProjectName + ` <command> [arguments]

Commands using kubectl:

	apply [package]     apply kubernetes package to cluster
	delete [package]    delete kubernetes package from cluster
	create [package]    create kubernetes package in cluster

Help command:

	help | --h          see help information
`
	fmt.Printf("%s", message)
	os.Exit(0)
}
