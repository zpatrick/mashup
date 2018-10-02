SHELL:=/bin/bash
VERSION?=$(shell git describe --tags --always)
CURRENT_DOCKER_IMAGE=zpatrick/mashup:$(VERSION)
LATEST_DOCKER_IMAGE=zpatrick/mashup:latest

run:
	go run main.go

generate:
	$(MAKE) -C generator run

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --ldflags "-X main.Version=$(VERSION)" -o mashup . 
	docker build -t $(CURRENT_DOCKER_IMAGE) .

release: build
	docker push $(CURRENT_DOCKER_IMAGE)
	docker tag  $(CURRENT_DOCKER_IMAGE) $(LATEST_DOCKER_IMAGE)
	docker push $(LATEST_DOCKER_IMAGE)

.PHONY: run generate build release
