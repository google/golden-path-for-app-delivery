FROM golang:1.10
WORKDIR /go/src/backend
COPY vendor/ ./vendor/
COPY *.go ./
RUN go build
CMD ["/go/src/backend/backend"]
