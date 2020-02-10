package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Build struct {
	Dockerfile []struct {
		Tag string `yaml:"tag"`
		Path string `yaml:"path"`
		Dockerfile string `yaml:"dockerfile"`
		Context string `yaml:"context"`
	}
}

/** Check file extension */
func CheckExtension(i os.FileInfo) bool {
	extensions := []string{".yaml", ".yml"}

	for _, ext := range extensions {
		if filepath.Ext(i.Name()) == ext {
			return true
		}
	}

	return false
}


/** Loop in the files */
func ReadFiles(dirname string, configParams Config) []string {
	dir := fmt.Sprintf("%s/%s", configParams.ConfigFolder, dirname)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var list []string

	for _, f := range files {
		// passing over dirs, or build files
		if f.IsDir() || f.Name() == "build.yaml" || f.Name() == "build.yml" {
			continue
		}

		// checking if extension is available
		if CheckExtension(f) {
			list = append(list, f.Name())
		}
	}

	return list
}


/** Checking if project exists */
func CheckIfProjectExists(config Config, dirName string, command string) bool {
	var dirs []os.FileInfo

	switch command {
	case "kube":
		dirs, _ = ioutil.ReadDir(config.ConfigFolder)
		break
	case "helm":
		dirs, _ = ioutil.ReadDir(config.HelmCharts)
		break
	default:
		Alert("ERR", "No such command configured!")
	}

	for _, dir := range dirs {
		if dir.Name() == dirName && dir.IsDir() {
			return true
		}
	}

	return false
}


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
			Alert("ERR", "Tag and path required!")
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
				Alert("ERR", "Already defined an tag in the `tag` definition from yaml!")
			}
			useTag = fmt.Sprintf("%s:%s", buildData.Tag, definedTag)
		} else {
			useTag = buildData.Tag
		}

		// create actual build command and execute
		cmd := fmt.Sprintf("docker build -t %s -f %s/%s %s", useTag, context,  dockerfile, context)

		fmt.Printf("\u001b[34m[NOTICE] Executing: \u001b[0m \u001b[36m%s \u001b[0m\n\n", cmd)
		ExecCommand(strings.Split(cmd, " "))
	}
}
