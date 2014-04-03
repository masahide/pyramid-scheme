.PHONY: all binary build default test 

default: binary

all: get test build

binary: get build

get:
	go get ./...

test: 
	go test -v 

build: 
	hack/make.sh
