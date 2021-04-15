ARG BASE
ARG SKAFFOLD_GO_GCFLAGS

FROM $BASE
COPY vendor/ vendor/
COPY templates/ *.go go.* ./
RUN go build -o app
