imageTag = viglesiasce/sample-app

all: build test push
PHONY: all

build:
	docker build -t $(imageTag) .

test:
	docker run $(imageTag) go test

push:
	gcloud docker -- push $(imageTag)
