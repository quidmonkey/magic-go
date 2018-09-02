export GOPATH := $(shell pwd)

build:
	cd src/magic-server && go build main.go

run:
	cd src/magic-server && go run main.go

setup:
	cd src/magic-server && dep ensure
