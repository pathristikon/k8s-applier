package utils

import (
	"fmt"
	"strings"
)


/** Handle the helm charts */
func HelmHandler(project string, cmd string, config Config) {
	cmdString := fmt.Sprintf("%s %s %s/%s", HelmArgument, cmd, config.HelmCharts,  project)
	command := strings.Split(cmdString, " ")

	Alert("NOTICE", fmt.Sprintf("Running: %s \n",  strings.Join(command, " ")), true )
	ExecCommand(command)
}