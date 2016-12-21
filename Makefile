imageRepo = viglesiasce/sample-app
versionFile = VERSION
majorVersion = $(shell awk 'BEGIN {FS="."}; {print $$1}' $(versionFile))
minorVersion = $(shell awk 'BEGIN {FS="."}; {print $$2}' $(versionFile))
subMinorVersion = $(shell awk 'BEGIN {FS="."}; {print $$3}' $(versionFile))
imageTag = v$(majorVersion).$(minorVersion)
image = $(imageRepo):$(imageTag)

all: build test push
PHONY: all

build:
	docker build -t $(image) .
	docker tag $(image) $(image).$(subMinorVersion)

test: build
	docker run $(image) go test

push:
	gcloud docker -- push $(image)
	gcloud docker -- push $(image).$(subMinorVersion)
	echo IMAGE=$(image) > image.properties
