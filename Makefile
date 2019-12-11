export GOPROXY=https://goproxy.io
default: build

build: export GO111MODULE=on

build:
	go build -o jiazhen main.go
