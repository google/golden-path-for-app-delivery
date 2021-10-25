# Copyright 2021 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.16

# Add a non-root user
RUN useradd -u 1000 -ms  /bin/bash app
RUN mkdir -p /go/src/app && chown -R app:app /go/src/app
USER app

# Cache dependencies
WORKDIR /go/src/app
COPY go.* ./
RUN go mod download

ENV GOTRACEBACK=all
ARG SKAFFOLD_GO_GCFLAGS
# Copy in source files
COPY *.go *.html ./
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o app
CMD ["/go/src/app/app"]
COPY k8s k8s
