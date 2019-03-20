# Build the manager binary
FROM golang:1.12-alpine as builder

# Copy in the go src
WORKDIR /go/src/github.com/dichque/grafana-operator
COPY pkg/    pkg/
COPY cmd/    cmd/
COPY vendor/ vendor/
COPY config/templates config

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager github.com/dichque/grafana-operator/cmd/manager
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o manager github.com/dichque/grafana-operator/cmd/manager

# Copy the controller-manager into a thin image
# FROM ubuntu:latest
FROM alpine
WORKDIR /
COPY --from=builder /go/src/github.com/dichque/grafana-operator/manager .
COPY --from=builder /go/src/github.com/dichque/grafana-operator/config config
ENTRYPOINT ["/manager"]
