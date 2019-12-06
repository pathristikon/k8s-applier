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
}


var configFolderString string


/** what to do if we just initialize the project */
func InitSystem() Config {
	currentUser, err := user.Current()

	if err != nil {
		panic("Cannot get current user!")
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
			panic("Couldn't create directory in Home folder!")
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

	questions := []string{"Where do you keep the projects?", "Where do you keep the k8s configs?"}
	answers := &Config{}

	for index, question := range questions {
		fmt.Println(question)
		answer, _ := reader.ReadString('\n')

		answer = strings.TrimSpace(answer)

		if index == 0 {
			answers.ProjectFolder = answer
		} else if index == 1 {
			answers.ConfigFolder = answer
		}
	}

	// encoding to json
	configJson, err := json.Marshal(answers)
	if err != nil {
		panic("Couldn't save the config to JSON")
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