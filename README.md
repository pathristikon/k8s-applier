# K8S Local cluster helper

### Installation

Run `make install-linux`

### Usage

`k8s <command> [arguments]`

### Configuration

If is the first time you execute the script, you will be asked to answer a few questions.
Please be sincere and provide the information requested.

The questions are:
    
    Where do you keep the projects?
    Where do you keep the k8s configs?

Provide absolute path.

### Help command

Run the following commands for help:

    k8s
    k8s --help
    k8s -help

### Commands with kubectl

These commands will apply/delete/create the yaml configuration yaml files for your
kubernetes cluster.

    k8s kubectl apply [package]
    k8s kubectl delete [package]
    k8s kubectl create [package]

The command is recursive and it will execute the command on all the yaml files found in
the package, except for an file called `build.yml` or `build.yaml`.

@author Dumitru Alexandru