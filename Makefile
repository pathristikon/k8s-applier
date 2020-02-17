.PHONY: install

PROJECT = k8s

install:
	cd src/utils && go test
	cd src/ && GO111MODULE=on go build -o ../k8s main.go

install-linux:
	cd src/ && GO111MODULE=on go build -o /usr/local/bin/${PROJECT} main.go