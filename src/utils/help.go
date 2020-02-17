package utils

import (
	"fmt"
	"os"
)


/** Help text */
func PrintHelp() {
	message := `Kubernetes cluster helper - v.` + K8sVersion + `

Usage:
	` + ProjectName + ` <flags> <command> [arguments]

Commands using kubectl:
	kube apply [package]     	apply kubernetes package to cluster
	kube delete [package]    	delete kubernetes package from cluster
	kube create [package]    	create kubernetes package in cluster

Commands using docker:
	build [package]          	build package based on yaml build.yml|yaml file

Commands using helm:
	helm install [package]   	install package from HelmCharts config folder
	helm uninstall [package] 	uninstalls package from HelmCharts config folder
	helm status [package]    	get status of package

Flags:
	-dry-run                    	Don't actually execute the command, 
					just print the messages 
	-push							Push image to docker registry after build
Help command:
	help | --h          		see help information
`
	fmt.Printf("%s", message)
	os.Exit(0)
}
