.PHONY: install

PROJECT = k8s

install:
	go build -o k8s src/main.go

install-linux:
	cd src/ && go build -o /usr/local/bin/${PROJECT} main.go