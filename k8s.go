package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"os/exec"
	"strings"
)

const ProjectDir = "./projects/"

func main() {
	arguments := checkArgs()

	project := arguments[0]

	yamlFiles :=readFiles(project)

	fmt.Printf("[Notice] Found %d files... continuing\n\n", len(yamlFiles))

	applyK8sYamls(yamlFiles)
}


/** apply the k8s yaml files */
func applyK8sYamls(files []string) {
	for _, file := range files {
		command := strings.Split("kubectl apply -f " + ProjectDir + file, " ")
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
		panic("Minimum one argument required. Format: [cmd] [folder] [options]")
	}

	if arguments[0] == "init" {
		initSystem()
		os.Exit(200)
	}

	dirs, _ := ioutil.ReadDir(ProjectDir)

	dirExists := false
	for _, dir := range dirs {
		if dir.Name() == arguments[0] && dir.IsDir() {
			dirExists = true
		}
	}

	if !dirExists {
		panic("Project doesn't exists!")
	}

	return arguments
}


/** what to do if we just initialize the project */
func initSystem() {
	fmt.Println("Initializing Kubernetes")
}


/** Loop in the files */
func readFiles(dirname string) []string {
	files, err := ioutil.ReadDir(ProjectDir + dirname)
	if err != nil {
		log.Fatal(err)
	}

	var list []string

	for _, f := range files {
		// passing over dirs
		if f.IsDir() {
			continue
		}

		// checking if extension is available
		if checkExtension(f) {
			list = append(list, f.Name())
		}
	}

	return list
}


/** Check file extension */
func checkExtension(i os.FileInfo) bool {
	extensions := []string{".yaml", ".yml"}

	for _, ext := range extensions {
		if filepath.Ext(i.Name()) == ext {
			return true
		}
	}

	return false
}
