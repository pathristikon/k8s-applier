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
	kube <global-args> apply [package]     	apply kubernetes package to cluster
	kube <global-args> delete [package]    	delete kubernetes package from cluster
	kube <global-args> create [package]    	create kubernetes package in cluster

Commands using docker:
	build <global-args> [package]          	build package based on yaml build.yml|yaml file

Commands using helm:
	helm <global-args> install [package]   	install package from HelmCharts config folder
	helm <global-args> uninstall [package] 	uninstalls package from HelmCharts config folder
	helm <global-args> status [package]    	get status of package

Global arguments:
	-dry-run                                Don't actually execute the command, 
                                            	just print the messages 

Help command:
	help | --h          see help information
`
	fmt.Printf("%s", message)
	os.Exit(0)
}
