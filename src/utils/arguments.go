package utils

import (
	"flag"
	"fmt"
	"os"
)


/** The global configuration struct */
type globalConfig struct {
	dryRun bool
	pushBuild bool
}


/** The global configuration for k8s-applier */
var appConfig globalConfig


/** Parsing the arguments from command line */
func ParseArguments() {
	/** Initializing system */
	configParams := InitSystem()

	/** Parse arguments */
	globalFlags()
	kubectl(configParams)
	dockerBuild(configParams)
	helm(configParams)

	/** Default behavior */
	Alert("ERR", "This command doesn't exists!", false)
	os.Exit(1)
}


/** Parse global flags */
func globalFlags() {
	help := flag.Bool("help", false, "Get help")
	h := flag.Bool("h", false, "Get help")
	dryRunFlag := flag.Bool("dry-run", false, "Dry run the commands without executing them")
	pushBuildFlag := flag.Bool("push", false, "Push docker image after building")
	flag.Parse()

	if *dryRunFlag {
		appConfig.dryRun = true
	}

	if *pushBuildFlag {
		appConfig.pushBuild = true
	}

	if len(os.Args) >= 2 && *help || len(os.Args) >= 2 && *h || len(os.Args) == 1 {
		PrintHelp()
	}
}

/** Basic flag parser for commands such as kubectl & helm */
func baseCommands(command string, configuredCommands map[string]bool, config Config) (string, string, Config, []string) {
	// escaping arguments that are globally defined
	escapeArgumentsAlreadyInConfig(os.Args)

	if os.Args[1] == command {
		parseArgs := flag.NewFlagSet(command, flag.ExitOnError)
		_ = parseArgs.Parse(os.Args[2:])

		args := parseArgs.Args()

		if len(args) < 2 {
			Alert("ERR", fmt.Sprintf("Expected %s [cmd] [project]", command), false)
		}

		cmd := args[0]
		project := args[1]

		otherArguments := args[2:]

		projectExists := CheckIfProjectExists(config, project, command)
		/** check if project folder exists */
		if !projectExists {
			Alert("ERR","Project folder does not exists!", false)
		}

		/** check if cmd is in map */
		if _, validChoice := configuredCommands[cmd]; !validChoice {
			Alert("ERR", "This kubernetes command can't be applied! Check help for details!", false)
		}

		return project, cmd, config, otherArguments
	}

	return "", "", Config{}, []string{}
}


/** Kubectl arguments */
func kubectl(config Config) {
	commands := map[string]bool{"apply": true, "delete": true, "create": true}
	project, cmd, config, _ := baseCommands(KubectlArgument, commands, config)
	if project != "" || cmd != "" {
		KubectlHandler(project, cmd, config)
		os.Exit(0)
	}
}


/** Helm arguments */
func helm(config Config) {
	commands := map[string]bool{"install": true, "delete": true, "status": true}
	project, cmd, config, otherArguments := baseCommands(HelmArgument, commands, config)

	otherArguments = escapeArgumentsAlreadyInConfig(otherArguments)

	if project != "" || cmd != "" {
	 	HelmHandler(project, cmd, config, otherArguments)
		os.Exit(0)
	}
}


/** Build dockerfiles based on YAML file arguments */
func dockerBuild(config Config) {
	var arv string
	if !appConfig.pushBuild {
		arv = os.Args[1]
	} else {
		arv = os.Args[2]
	}

	if arv == "build" {
		commands := flag.NewFlagSet("build", flag.ExitOnError)
		tag := commands.String("tag", "", "Choose the tag of the docker image being built")
		if !appConfig.pushBuild {
			_ = commands.Parse(os.Args[2:])
		} else {
			_ = commands.Parse(os.Args[3:])
		}

		args := commands.Args()
		if len(args) < 1 {
			Alert("ERR","Expected build [project]", false)
		}

		BuildDockerImages(config, args[0], *tag)
		os.Exit(0)
	}
}
