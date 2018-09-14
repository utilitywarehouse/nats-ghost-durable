FROM golang:1.11-alpine

ADD . /go/src/github.com/utilitywarehouse/nats-ghost-durable/sub
WORKDIR /go/src/github.com/utilitywarehouse/nats-ghost-durable/sub

RUN apk --no-cache add git ca-certificates && \
  go get -v -d ./... && \
  go build && \
  mv sub /sub

ENTRYPOINT ["/sub"]