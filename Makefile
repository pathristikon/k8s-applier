.PHONY: install

PROJECT = k8s

install:
	cd src/ && go build -o k8s main.go

install-linux:
	cd src/ && go build -o /usr/local/bin/${PROJECT} main.go