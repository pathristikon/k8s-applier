package main

import (
	"./utils"
	"fmt"
	"io/ioutil"
	"os"
)

var configParams utils.Config

// [cmd] apply [project]
func main() {
	configParams = utils.InitSystem()
	arguments := checkArgs()
	project := arguments[1]

	// apply or delete
	cmd := arguments[0]

	yamlFiles := utils.ReadFiles(project, configParams)

	fmt.Printf("[Notice] Found %d files... continuing\n\n", len(yamlFiles))

	utils.ApplyK8sYamls(yamlFiles, cmd, project, configParams)
}


/** get arguments */
func checkArgs() []string {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		utils.PrintHelp()
	}

	if arguments[0] == "--help" || arguments[0] == "-h" {
		utils.PrintHelp()
	}

	// so here we need the project name
	if len(arguments) < 2 {
		panic("Valid format: [cmd] apply/delete [project]")
	}

	dirs, _ := ioutil.ReadDir(configParams.ConfigFolder)

	dirExists := false
	for _, dir := range dirs {
		if dir.Name() == arguments[1] && dir.IsDir() {
			dirExists = true
		}
	}

	if !dirExists {
		panic("Project doesn't exists!")
	}

	return arguments
}
