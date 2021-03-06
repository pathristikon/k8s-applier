package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)


type Config struct {
	ProjectFolder string `json:"projectFolder,omitempty"`
	ConfigFolder  string `json:"configFolder,omitempty"`
	HelmCharts  string `json:"helmCharts,omitempty"`
}


var configFolderString string


/** what to do if we just initialize the project */
func InitSystem() Config {
	currentUser, err := user.Current()

	if err != nil {
		Alert("ERR", "Cannot get current user!", false)
	}

	configFolderString = fmt.Sprintf("%s/%s", currentUser.HomeDir, ConfigDir)

	// check if home directory contains config folder
	if createConfigFolderIfNotExists(currentUser) {
		resolveQuestions()
	}

	return getConfigStruct()
}


/** create config folder if not exists */
func createConfigFolderIfNotExists(currentUser *user.User) bool {
	isConfigPresent := getConfigFolder(currentUser)

	if !isConfigPresent {
		fmt.Println("Config not present! Creating config folder...")
		err := os.Mkdir(configFolderString, 0755)

		if err != nil {
			Alert("ERR", "Couldn't create directory in Home folder!", false)
		}

		return true
	}

	return false
}


/** check if config folder exists */
func getConfigFolder(currentUser *user.User) bool {
	dirs, err := ioutil.ReadDir(currentUser.HomeDir)
	if err != nil {
		panic(err)
	}

	for _, d := range dirs {
		if d.IsDir() && d.Name() == ConfigDir {
			return true
		}
	}

	return false
}


/** saving the config to json file */
func resolveQuestions() {
	reader := bufio.NewReader(os.Stdin)

	questions := []string{
		"Where do you keep the projects?",
		"Where do you keep the k8s configs?",
		"[Optional] Where do you keep the Helm charts?",
	}

	answers := &Config{}

	for index, question := range questions {
		var answer string

		for {
			fmt.Println(question)
			answer, _ = reader.ReadString('\n')
			answer = strings.TrimSpace(answer)

			// skip optional parameter Config.HelmCharts
			if index == 2 || len(answer) > 0 {
				break
			}
		}

		switch index {
		case 0:
			answers.ProjectFolder = answer
			break
		case 1:
			answers.ConfigFolder = answer
			break
		case 2:
			answers.HelmCharts = answer
			break
		}
	}

	// encoding to json
	configJson, err := json.Marshal(answers)
	if err != nil {
		Alert("ERR", "Couldn't save the config to JSON", false)
	}

	filename := fmt.Sprintf("%s/%s", configFolderString, ConfigFile)
	var file, fileError = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

	if fileError != nil {
		panic(fileError.Error())
	}

	defer file.Close()

	_, errWrite := file.WriteString(string(configJson))
	if errWrite != nil {
		panic(errWrite.Error())
	}
	_ = file.Sync()

	fmt.Println("\nWrote config successfully!")
	fmt.Println(string(configJson))
}


/** retrieve config from json to struct */
func getConfigStruct() Config {
	fileName := fmt.Sprintf("%s/%s", configFolderString, ConfigFile)
	configFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}

	bytes := []byte(string(configFile))
	var config Config

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		panic(err.Error())
	}

	return config
}
