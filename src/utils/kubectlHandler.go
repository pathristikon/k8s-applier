package utils

import "fmt"


/** Handle the kubectl yaml files */
func KubectlHandler(project string, cmd string, config Config) {
	yamlFiles := ReadFiles(project, config)

	fmt.Printf("[Notice] Found %d files... continuing\n\n", len(yamlFiles))

	HandleYamls(yamlFiles, cmd, project, config)
}
