package main

import (
	"./utils"
	"fmt"
	"io/ioutil"
	"os"
)


// [cmd] apply [project]
func main() {
	arguments := checkArgs()

	project := arguments[1]

	// apply or delete
	cmd := arguments[0]

	yamlFiles := utils.ReadFiles(project)

	fmt.Printf("[Notice] Found %d files... continuing\n\n", len(yamlFiles))

	utils.ApplyK8sYamls(yamlFiles, cmd, project)
}


/** get arguments */
func checkArgs() []string {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		panic("Valid format: [cmd] apply/delete [project]")
	}

	if arguments[0] == "init" {
		utils.InitSystem()
		os.Exit(200)
	}

	// so here we need the project name
	if len(arguments) < 2 {
		panic("Valid format: [cmd] apply/delete [project]")
	}

	dirs, _ := ioutil.ReadDir(utils.ProjectDir)

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
