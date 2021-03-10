#!/bin/bash -xe
gcloud beta artifacts docker images scan $1
