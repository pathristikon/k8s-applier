.PHONY: install

PROJECT = k8s

install:
	cd src/ && go test ./utils && GO111MODULE=on go build -o ../k8s main.go

install-linux:
	cd src/ && GO111MODULE=on go build -o /usr/local/bin/${PROJECT} main.go