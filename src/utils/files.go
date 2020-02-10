package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
	case KubectlArgument:
		dirs, _ = ioutil.ReadDir(config.ConfigFolder)
		break
	case HelmArgument:
		dirs, _ = ioutil.ReadDir(config.HelmCharts)
		break
	default:
		Alert("ERR", "No such command configured!", false)
	}

	for _, dir := range dirs {
		if dir.Name() == dirName && dir.IsDir() {
			return true
		}
	}

	return false
}
