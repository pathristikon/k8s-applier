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

	kube apply [package]     apply kubernetes package to cluster
	kube delete [package]    delete kubernetes package from cluster
	kube create [package]    create kubernetes package in cluster

Help command:

	help | --h          see help information
`
	fmt.Printf("%s", message)
	os.Exit(0)
}
