GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
LDFLAGS := -s -w
GOARCH=amd64
SERVICE=tron-scan
MAIN=tron.go

.PHONY: build-linux
build-linux:
	env CGO_ENABLED=0 GOOS=linux GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -trimpath -o $(SERVICE) $(MAIN)

.PHONY: build-win
build-win:
	env CGO_ENABLED=0 GOOS=windows GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -trimpath -o $(SERVICE).exe $(MAIN)

.PHONY: build-mac
build-mac:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -trimpath -o $(SERVICE) $(MAIN)

.PHONY: build-docker
build-docker:
	docker build -t ${DOCKER_USERNAME}/$(SERVICE):${VERSION} .
	@echo "Build docker successfully"

.PHONY: publish-docker
publish-docker:
	echo "${DOCKER_PASSWORD}" | docker login --username ${DOCKER_USERNAME} --password-stdin
	docker push ${DOCKER_USERNAME}/$(SERVICE):${VERSION}
	@echo "Publish docker successfully"
