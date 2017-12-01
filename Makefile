DOCKER_REGISTRY ?= "bacongobbler"

.PHONY: build
build:
	mkdir -p bin/
	go build -o bin/github-notify ./main.go

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_REGISTRY)/github-notify:latest .

.PHONY: docker-push
docker-push:
	docker push $(DOCKER_REGISTRY)/github-notify

.PHONY: bootstrap
bootstrap:
	dep ensure
