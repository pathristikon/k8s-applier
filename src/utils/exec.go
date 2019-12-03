package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

/** apply the k8s yaml files */
func ApplyK8sYamls(files []string, cmd string, project string, configParams Config) {
	for _, file := range files {
		command := strings.Split("kubectl " + cmd + " -f " + configParams.ConfigFolder + "/" + project + "/" + file, " ")

		fmt.Printf("Running: %s \n",  strings.Join(command, " "))

		ExecCommand(command)
	}
}


/** execute command and print response */
func ExecCommand(args []string) {
	cmd := exec.Command(args[0], args[1:]...)

	stdout, stderr := cmd.CombinedOutput()

	if stderr != nil {
		fmt.Println("[ERR] Cannot execute command!")
	}

	println(string(stdout))
}
