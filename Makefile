PROJECT = k8s

install-linux:
	go build -o /usr/local/bin/${PROJECT} src/main.go