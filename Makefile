export GOPROXY=https://goproxy.io
default: build

build_run: build
	./jiazhen

build: export GO111MODULE=on

build:
	swag init
	go build -o jiazhen main.go


