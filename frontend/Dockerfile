FROM golang:1.16

# Add a non-root user
RUN useradd -u 1000 -ms  /bin/bash app
WORKDIR /go/src/app
RUN chown -R app:app /go/src/app
USER app

# Copy in source files
COPY vendor/ vendor/
COPY templates/ templates/
COPY *.go go.* ./

ENV GOTRACEBACK=single
ARG SKAFFOLD_GO_GCFLAGS
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o app

CMD ["/go/src/app/app"]