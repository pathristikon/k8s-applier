package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

/** apply the k8s yaml files */
func ApplyK8sYamls(files []string, cmd string, project string, configParams Config) {
	for _, file := range files {
		cmdString := fmt.Sprintf("kubectl %s -f %s/%s/%s", cmd, configParams.ConfigFolder, project, file)
		command := strings.Split(cmdString, " ")

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

	fmt.Println(string(stdout))
}
