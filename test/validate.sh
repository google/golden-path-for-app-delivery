#!/bin/bash -xe
cd frontend
kustomize build k8s/prod > ../prod.yaml
cd ..
conftest test -p policy ./prod.yaml
