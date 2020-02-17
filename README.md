[![Build Status](https://travis-ci.org/pathristikon/k8s-applier.svg?branch=master)](https://travis-ci.org/pathristikon/k8s-applier)

# K8S Local cluster helper

### Installation

Run `make install-linux`

### Usage

`k8s <flags> <command> [arguments]`

### Configuration

If is the first time you execute the script, you will be asked to answer a few questions.
Please be sincere and provide the information requested.

The questions are:
    
    Where do you keep the projects?
    Where do you keep the k8s configs?
    [Optional] Where do you keep the Helm charts?

Provide absolute path.

### Help command

Run the following commands for help:

    k8s
    k8s --help
    k8s -help

### Commands using kubectl

These commands will apply/delete/create the yaml configuration yaml files for your
kubernetes cluster.

    k8s kubectl apply [package]
    k8s kubectl delete [package]
    k8s kubectl create [package]

The command is recursive and it will execute the command on all the yaml files found in
the package, except for an file called `build.yml` or `build.yaml`.

### Commands using helm
    k8s helm install [package]
    k8s helm uninstall [package]
    k8s helm status [package]
    
If you would like to pass extra arguments, feel free to do it so:
    `k8s helm install [package] --generate-name`

### Commands using docker
    k8s build [package] OR
    k8s build -tag [tag] [package]
    
The build package will build your Dockerfiles based on specifications from build.yml or build.yaml.

Minimum required format build.yml file:

```json
dockerfile:
    - tag: "mycoolapp:latest"
      path: "infrastructure/mycoolapp"
    
    - tag: "mycoolappdb:3fr5310"
      path: "infrastructure/mycoolapp/db"
      dockerfile: "DB_Dockerfile"
```

### Flags
Flags are:
    
    -dry-run       Don't actually execute the command, just print the messages 

#### Options:
- tag: the tag used to build your docker image
- path: the path to docker context. If is not prefixed with "/", it will be relative to
your projects' folder from config file.
- dockerfile: name of the dockerfile, if multiple names - optional. Default: `Dockerfile`

#### Tag specifications:

A tag can be set either in YAML file `mycoolapp:latest` where latest is the tag or using the 
`-- tag` flag.

The tag will apply to all docker images created.


@author Dumitru Alexandru