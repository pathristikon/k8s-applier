.PHONY: install

PROJECT = k8s

install:
	go build src/main.go

install-linux:
	go build -o /usr/local/bin/${PROJECT} src/main.go