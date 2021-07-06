FROM golang:1.16

# Add a non-root user
RUN useradd -u 1000 -ms  /bin/bash app
WORKDIR /go/src/app
RUN chown -R app:app /go/src/app
USER app

# Copy in source files
COPY vendor/ ./vendor/
COPY *.go go.* *.html ./

ENV GOTRACEBACK=single
ARG SKAFFOLD_GO_GCFLAGS
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o app
COPY k8s k8s
CMD ["/go/src/app/app"]