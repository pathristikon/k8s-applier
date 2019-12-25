package utils

import (
	"flag"
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


/** Kubectl arguments */
func kubectl(config Config) {
	if os.Args[1] == "kube" {
		kube := flag.NewFlagSet("kube", flag.ExitOnError)
		_ = kube.Parse(os.Args[2:])

		args := kube.Args()
		if len(args) < 2 {
			Alert("ERR","Expected kube [cmd] [project]")
		}

		cmd := args[0]
		project := args[1]
		projectExists := CheckIfProjectExists(config, project)

		/** check if cmd is in map */
		choices := map[string]bool{"apply": true, "delete": true, "create": true}
		if _, validChoice := choices[cmd]; !validChoice {
			Alert("ERR", "This kubernetes command can't be applied! Check help for details!")
		}

		/** check if project folder exists */
		if !projectExists {
			Alert("ERR","Project folder does not exists!")
		}

		HandleKubernetesFiles(project, cmd, config)
		os.Exit(0)
	}
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