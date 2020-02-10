package utils

import (
	"fmt"
	"strings"
)


/** Handle the helm charts */
func HelmHandler(project string, cmd string, config Config, otherArguments []string) {
	addOtherArguments := strings.Join(otherArguments, " ")
	cmdString := fmt.Sprintf("%s %s %s/%s %s", HelmArgument, cmd, config.HelmCharts,  project, addOtherArguments)
	command := strings.Split(cmdString, " ")

	Alert("NOTICE", fmt.Sprintf("Running: %s \n",  strings.Join(command, " ")), true )
	ExecCommand(command)
}