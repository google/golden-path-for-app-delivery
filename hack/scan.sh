#!/bin/bash -xe
export SCAN=$(gcloud beta artifacts docker images scan $1 --format json | jq -r .response.scan)
gcloud beta artifacts docker images list-vulnerabilities ${SCAN} > hack/scan.yaml
