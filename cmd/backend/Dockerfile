FROM golang:1.16
RUN useradd -u 1000 -ms  /bin/bash app
WORKDIR /go/src/backend
RUN chown -R app:app /go/src/backend
USER app

COPY vendor/ ./vendor/
COPY *.go go.* ./
RUN go build
CMD ["/go/src/backend/backend"]
