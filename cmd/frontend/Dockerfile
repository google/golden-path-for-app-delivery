FROM golang:1.16
RUN useradd -u 1000 -ms  /bin/bash app
WORKDIR /go/src/frontend
RUN chown -R app:app /go/src/frontend
USER app

COPY *.go go.* ./
RUN go build
CMD ["/go/src/frontend/frontend"]
