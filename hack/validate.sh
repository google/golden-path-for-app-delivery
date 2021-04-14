#!/bin/bash -xe

docker run -v `pwd`:/src $1 bash -c 'kustomize build /src/k8s/dev > /src/dev.yaml'
docker run -v `pwd`:/src $1 bash -c 'conftest test -p /src/hack/policy /src/dev.yaml'