package main

import (
	"./utils"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)


// [cmd] apply [project]
func main() {
	arguments := checkArgs()

	project := arguments[1]

	// apply or delete
	cmd := arguments[0]

	yamlFiles := utils.ReadFiles(project)

	fmt.Printf("[Notice] Found %d files... continuing\n\n", len(yamlFiles))

	applyK8sYamls(yamlFiles, cmd, project)
}


/** apply the k8s yaml files */
func applyK8sYamls(files []string, cmd string, project string) {
	for _, file := range files {
		command := strings.Split("kubectl " + cmd + " -f " + utils.ProjectDir + project + "/" + file, " ")

		fmt.Printf("Running: %s \n",  strings.Join(command, " "))

		execCommand(command)
	}
}


/** execute command and print response */
func execCommand(args []string) {
	cmd := exec.Command(args[0], args[1:]...)

	stdout, stderr := cmd.CombinedOutput()

	if stderr != nil {
		panic(stderr.Error())
	}

	println(string(stdout))
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
