package utils

import (
	"fmt"
	"os"
)


/** Help text */
func PrintHelp() {
	message := `Kubernetes cluster helper

Usage:
	` + ProjectName + ` <command> [arguments]

Commands using kubectl:
	kube apply [package]     apply kubernetes package to cluster
	kube delete [package]    delete kubernetes package from cluster
	kube create [package]    create kubernetes package in cluster

Commands using docker:
	build [package]          build package based on yaml build.yml|yaml file

Commands using helm:
	helm install [package]   install package from HelmCharts config folder

Help command:
	help | --h          see help information
`
	fmt.Printf("%s", message)
	os.Exit(0)
}
