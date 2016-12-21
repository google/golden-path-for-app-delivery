imageRepo = viglesiasce/sample-app
imageTag = v0.1.0
image = $(imageRepo):$(imageTag)

all: build test push
PHONY: all

build:
	docker build -t $(image) .

test:
	docker run $(image) go test

push:
	gcloud docker -- push $(image)
	echo IMAGE=$(image) > image.properties
