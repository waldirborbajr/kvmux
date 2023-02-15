LAST_TAG=$(shell git describe --abbrev=0 --tags)
CURR_SHA=$(shell git rev-parse --verify HEAD)

LDFLAGS=-ldflags "-s -w -X main.version=$(LAST_TAG)"

release:
	git tag $(tag)
	git push origin $(tag)

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-drpc_out=. --go-drpc_opt=paths=source_relative --proto_path=. actor/actor.proto

test: build
	go test ./... --race
	
build:
	GOOS=$(os) GOARCH=amd64 go build ${LDFLAGS} -o bin/kvmux cmd/main.go

install:
	cp bin/kvmux ${HOME}/.local/bin

clean:
	rm bin/kvmux
	rm ${HOME}/.local/bin/kvmux

.PHONY: proto