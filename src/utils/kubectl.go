package utils

import "fmt"

func HandleKubernetesFiles(project string, cmd string, config Config) {
	yamlFiles := ReadFiles(project, config)

	fmt.Printf("[Notice] Found %d files... continuing\n\n", len(yamlFiles))

	HandleYamls(yamlFiles, cmd, project, config)
}
