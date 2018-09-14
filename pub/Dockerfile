FROM golang:1.11-alpine

ADD . /go/src/github.com/utilitywarehouse/nats-ghost-durable/pub
WORKDIR /go/src/github.com/utilitywarehouse/nats-ghost-durable/pub

RUN apk --no-cache add git ca-certificates && \
  go get -v -d ./... && \
  go build && \
  mv pub /pub

ENTRYPOINT ["/pub"]