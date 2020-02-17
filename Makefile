.PHONY: travis

PROJECT = k8s

travis:
	cd ${TRAVIS_BUILD_DIR}/src/ && go test ./utils && GO111MODULE=on go build -o ../k8s main.go

install-linux:
	cd src/ && GO111MODULE=on go build -o /usr/local/bin/${PROJECT} main.go