#!/bin/bash -xe
if [[ $(kubectl config current-context) == "minikube1" ]];then
  eval $(minikube -p minikube docker-env)
fi

export SCAN=$(gcloud beta artifacts docker images scan $1 --format json | jq -r .response.scan)
gcloud beta artifacts docker images list-vulnerabilities ${SCAN}
