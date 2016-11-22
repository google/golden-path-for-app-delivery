#!/bin/bash
IMAGE_TAG=sample-app-$BUILD_NUMBER
docker build -t sample-app-$BUILD_NUMBER .
docker run $IMAGE_TAG go test
