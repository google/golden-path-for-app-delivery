#!/bin/bash -xe
kustomize build src/frontend/k8s/prod > frontend-prod.yaml
kustomize build src/backend/k8s/prod >  backend-prod.yaml
conftest test -p policy ./*-prod.yaml
