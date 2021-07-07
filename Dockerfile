FROM golang:1.16

# Add a non-root user
RUN useradd -u 1000 -ms  /bin/bash app
RUN mkdir -p /go/src/app && chown -R app:app /go/src/app
USER app

# Copy in source files
WORKDIR /go/src/app
COPY vendor/ ./vendor/
COPY *.go go.* *.html ./

ENV GOTRACEBACK=all
ARG SKAFFOLD_GO_GCFLAGS
RUN go build -trimpath -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o app
CMD ["/go/src/app/app"]
COPY k8s k8s
