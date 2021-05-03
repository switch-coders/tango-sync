GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/tango-sync
PACKAGES_PATH = $(shell go list -f '{{ .Dir }}' ./...)

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

clean:
	rm -rf $(DOCKER_BUILD)

heroku: $(DOCKER_CMD)
	heroku container:push web

.PHONY: gci
gci:
	@echo "Executing gci"
	@gci -w -local github.com/switch-coders/tango-sync $(PACKAGES_PATH)