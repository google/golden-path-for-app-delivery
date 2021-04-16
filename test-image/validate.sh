#!/bin/bash -xe
kustomize build cmd/frontend/k8s/prod > frontend-prod.yaml
kustomize build cmd/backend/k8s/prod >  backend-prod.yaml
conftest test -p test-image/policy ./*-prod.yaml
