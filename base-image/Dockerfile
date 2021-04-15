FROM golang:1.16

# Set up non-root user
RUN useradd -u 1000 -ms  /bin/bash app
WORKDIR /go/src/app
RUN chown -R app:app /go/src/app
USER app

# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=single
CMD ["/go/src/app/app"]