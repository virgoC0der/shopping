PROJECTNAME=$(shell basename "$(PWD)")

build:
	export GO111MODULE=on
	export GOPROXY=https://goproxy.cn,direct
	export GOROOT=/usr/local/Cellar/go/1.17.2/libexec
	go mod tidy
	go build -o $(PROJECTNAME)
