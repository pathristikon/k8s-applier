package utils

import (
	"flag"
	"fmt"
	"os"
)


/** Parsing the arguments from command line */
func ParseArguments() {
	/** Initializing system */
	configParams := InitSystem()

	/** Parse arguments */
	help()
	kubectl(configParams)
	dockerBuild(configParams)
	helm(configParams)

	/** Default behavior */
	Alert("ERR", "This command doesn't exists!")
	os.Exit(1)
}


/** Print help */
func help() {
	help := flag.Bool("help", false, "Get help")
	flag.Parse()

	if len(os.Args) >= 2 && *help || len(os.Args) == 1 {
		PrintHelp()
	}
}

func baseCommands(command string, configuredCommands map[string]bool, config Config) (string, string, Config) {
	if os.Args[1] == command {
		parseArgs := flag.NewFlagSet(command, flag.ExitOnError)
		_ = parseArgs.Parse(os.Args[2:])

		args := parseArgs.Args()
		if len(args) < 2 {
			Alert("ERR", fmt.Sprintf("Expected %s [cmd] [project]", command))
		}

		cmd := args[0]
		project := args[1]

		projectExists := CheckIfProjectExists(config, project, command)
		/** check if project folder exists */
		if !projectExists {
			Alert("ERR","Project folder does not exists!")
		}

		/** check if cmd is in map */
		if _, validChoice := configuredCommands[cmd]; !validChoice {
			Alert("ERR", "This kubernetes command can't be applied! Check help for details!")
		}

		return project, cmd, config
	}

	return "", "", Config{}
}

/** Kubectl arguments */
func kubectl(config Config) {
	commands := map[string]bool{"apply": true, "delete": true, "create": true}
	project, cmd, config := baseCommands("kubectl", commands, config)
	HandleKubernetesFiles(project, cmd, config)
	os.Exit(0)
}

/** Helm arguments */
func helm(config Config) {
	commands := map[string]bool{"install": true, "uninstall": true, "create": true}
	project, cmd, config := baseCommands("helm", commands, config)
	HandleKubernetesFiles(project, cmd, config)
	os.Exit(0)
}


/** Build dockerfiles based on YAML file arguments */
func dockerBuild(config Config) {
	if os.Args[1] == "build" {
		commands := flag.NewFlagSet("build", flag.ExitOnError)
		tag := commands.String("tag", "", "Choose the tag of the docker image being built")

		_ = commands.Parse(os.Args[2:])

		args := commands.Args()
		if len(args) < 1 {
			Alert("ERR","Expected build [project]")
		}

		BuildDockerImages(config, args[0], *tag)
		os.Exit(0)
	}
}