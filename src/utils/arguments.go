package utils

import (
	"flag"
	"os"
)

func ParseArguments() {
	/** Initializing system */
	configParams := InitSystem()

	/** Parse arguments */
	help()
	kubectl(configParams)

	/** Default behavior */
	Alert("ERR", "This command doesn't exists!")
	os.Exit(1)
}


func help() {
	help := flag.Bool("help", false, "Get help")
	flag.Parse()

	if len(os.Args) >= 2 && *help || len(os.Args) == 1 {
		PrintHelp()
	}
}


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
	}

	return
}
