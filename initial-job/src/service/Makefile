export GOPATH := $(shell pwd)/../../
export PATH := ${GOPATH}bin:${PATH}
export DEP_BIN := ${GOPATH}bin/
export GO111MODULE=on

env:
	 cp -n env.example env.local

dep-install:
	 ${DEP_BIN}dep ensure -v

build:
	 go build -o ./out/app ./src/main.go
	 chmod +x ./out/app

run:
	 ./out/app -log_dir=\"/tmp\" -alsologtostderr=true -logtostderr=true -v=2
