package utils

import (
	"fmt"
	"strings"
)


/** Handle the helm charts */
func HelmHandler(project string, cmd string, config Config, otherArguments []string) {
	var parsedCommand string

	addOtherArguments := strings.Join(otherArguments, " ")
	if cmd == "install" {
		parsedCommand = fmt.Sprintf("%s %s %s %s/%s/ %s",
			HelmArgument, cmd, project, config.HelmCharts,  project, addOtherArguments)
	} else if cmd == "delete" || cmd == "status" {
		parsedCommand = fmt.Sprintf("%s %s $(helm ls | grep %s | awk '{print $1}')",
			HelmArgument, cmd, project)
	}

	command := strings.Split(parsedCommand, " ")

	Alert("NOTICE", fmt.Sprintf("Running: %s \n",  strings.Join(command, " ")), true )

	if !appConfig.dryRun {
		ExecCommand(command)
	}
}
