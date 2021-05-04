.PHONY: default
default: help ;

coverage-console-local:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

coverage-console-docker:
	docker run --rm -v "$(PWD)":/usr/src/authorize -w /usr/src/authorize golang:1.16.3 go test ./... -coverprofile=coverage.out && go tool cover -func=coverage.out

coverage-html-local:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

coverage-html-docker:
	docker run --rm -v "$(PWD)":/usr/src/authorize -w /usr/src/authorize golang:1.16.3 go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

coverage-console: coverage-console-docker ;

coverage-html: coverage-html-docker ;

build-linux:
	rm -f bin/authorize-linux-amd64
	export DOCKER_BUILDKIT=1
	docker build --target bin --platform linux/amd64 --output bin/ .
	mv bin/authorize bin/authorize-linux-amd64

build-darwin:
	rm -f bin/authorize-darwin-amd64
	export DOCKER_BUILDKIT=1
	docker build --target bin --platform darwin/amd64 --output bin/ .
	mv bin/authorize bin/authorize-darwin-amd64

help:
	@echo
	@echo "Commands available:"
	@echo "  make coverage-console		Run unit tests with coverage output in terminal"
	@echo "  make coverage-html		Run unit tests with coverage output in browser"
	@echo "  make build-linux		Build project binary for linux with output at bin folder (./bin/authorize-linux-amd64)"
	@echo "  make build-darwin		Build project binary for macOS with output at bin folder (./bin/authorize-darwin-amd64)"
	@echo "  make help			Print this help"
	@echo
