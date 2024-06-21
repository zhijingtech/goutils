GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
GO_BUILD_EXE:=go
GO_BUILD_ARGS:=-ldflags "-X main.Version=$(VERSION)"

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
endif

# 打印所有 make 指令
help:
	@echo ''
	@echo '用法:'
	@echo ' make [选项]'
	@echo ''
	@echo '所有选项:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

.PHONY: all
# 执行：pb gen test
all: fmt lint test

.PHONY: init
# 安装格式化、Lint等工具
init:
	go install mvdan.cc/gofumpt@latest
	go install golang.org/x/tools/cmd/goimports@latest
	# go install github.com/incu6us/goimports-reviser/v3@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: fmt
# 格式化代码
fmt: 
	gofmt -w -r 'interface{} -> any' . 2>&1
	gofumpt -w -l . 2>&1
	# golangci-lint run --fix ./...
	go mod tidy

.PHONY: lint
# 代码静态检查
lint:
	golangci-lint run ./...

# .PHONY: build
# # 编译持续
# build: clean
# 	mkdir -p build/
# 	$(GO_BUILD_EXE) build $(GO_BUILD_ARGS) -o ./build/ .

.PHONY: test
# 执行所有单测
test:
	go test -coverprofile=profile.cov ./...


# .PHONY: clean
# # 清理编译生成的文件
# clean:
# 	rm -rf ./build
# 	rm -rf ./logs
# 	rm -f profile.cov