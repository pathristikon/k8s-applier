package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

/** apply the k8s yaml files */
func HandleYamls(files []string, cmd string, project string, configParams Config) {
	for _, file := range files {
		cmdString := fmt.Sprintf("%s %s -f %s/%s/%s", KubectlArgument, cmd, configParams.ConfigFolder, project, file)
		command := strings.Split(cmdString, " ")

		Alert("NOTICE", fmt.Sprintf("Running: %s \n",  strings.Join(command, " ")), true )
		ExecCommand(command)
	}
}


/** execute command and print response */
func ExecCommand(args []string) {
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout, cmd.Stderr = mw, mw

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		Alert("ERR", "Cannot execute command!", false)
	}

	fmt.Println(stdBuffer.String())
}
