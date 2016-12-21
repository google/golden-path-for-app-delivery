imageRepo = viglesiasce/sample-app
imageTag = v0.1
minorVersion = 1
image = $(imageRepo):$(imageTag)

all: build test push
PHONY: all

build:
	docker build -t $(image) .
	docker build -t $(image).$(minorVersion) .

test:
	docker run $(image) go test

push:
	gcloud docker -- push $(image)
	gcloud docker -- push $(image).$(minorVersion)
	echo IMAGE=$(image) > image.properties
