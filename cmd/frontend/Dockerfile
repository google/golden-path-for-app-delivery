ARG BASE
ARG SKAFFOLD_GO_GCFLAGS

FROM $BASE
COPY vendor/ vendor/
COPY templates/ templates/
COPY *.go go.* ./
RUN go build -o app
