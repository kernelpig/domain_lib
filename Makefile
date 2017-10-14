all: build

export GOPATH := $(GOPATH)
export GOBIN := $(CURDIR)/bin

CURRENT_GIT_GROUP := wangqingang
CURRENT_GIT_REPO := domain_lib

build:
	protoc -I pb pb/domain.proto --go_out=pb

test:
	go test -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/domain

clean:
	@rm -rf bin

.PHONY: test build clean
