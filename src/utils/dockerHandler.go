package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)


/** Build Dockerfile files based on build.yaml|yml */
func BuildDockerImages(config Config, project string, definedTag string) {

	var build Build

	filename := fmt.Sprintf("%s/%s/%s", config.ConfigFolder, project, "build.yaml")
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		panic("Cannot read build.yaml|yml file")
	}

	err = yaml.Unmarshal(file, &build)

	if err != nil {
		panic(fmt.Sprintf("Cannot parse file %s", filename))
	}

	for _, buildData := range build.Dockerfile {
		// check if the required arguments are set
		if buildData.Path == "" || buildData.Tag == "" {
			Alert("ERR", "Tag and path required!", false)
		}

		var context string
		var dockerfile string

		// setting up the name of the dockerfile
		if buildData.Dockerfile == "" {
			dockerfile = "Dockerfile"
		} else {
			dockerfile = buildData.Dockerfile
		}

		// checking if string has prefix "/"
		if strings.HasPrefix(buildData.Path, "/") {
			context = buildData.Path
		} else {
			context = fmt.Sprintf("%s/%s", config.ProjectFolder, buildData.Path)
		}

		// checking if tag is set and definition from yaml meets the criteria
		var useTag string
		if definedTag != "" {
			if strings.Contains(buildData.Tag, ":") {
				Alert("ERR", "Already defined an tag in the `tag` definition from yaml!", false)
			}
			useTag = fmt.Sprintf("%s:%s", buildData.Tag, definedTag)
		} else {
			useTag = buildData.Tag
		}

		// create actual build command and execute
		cmd := fmt.Sprintf("docker build -t %s -f %s/%s %s", useTag, context,  dockerfile, context)
		push := fmt.Sprintf("docker push %s", useTag)
		Alert("NOTICE", "Executing: " + cmd, true)

		if appConfig.pushBuild {
			Alert("NOTICE", "Executing: " + push, true)
		}

		if !appConfig.dryRun {
			ExecCommand(strings.Split(cmd, " "))
			if appConfig.pushBuild {
				ExecCommand(strings.Split(push, " "))
			}
		}
	}
}
