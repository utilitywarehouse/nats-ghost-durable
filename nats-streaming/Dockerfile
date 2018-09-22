FROM alpine:latest

ENV GOPATH /go

RUN apk --no-cache add gcc musl-dev git ca-certificates go && \
  mkdir /go && \
  go get github.com/nats-io/nats-streaming-server && \
  cd /go/src/github.com/nats-io/nats-streaming-server && \
  git checkout fix_ghost_durables && \
  go build . && \
  mv nats-streaming-server /

EXPOSE 4222 8222

ENTRYPOINT ["/nats-streaming-server"]
CMD ["-m", "8222", "--store=FILE", "--dir=/data", "--max_channels=0", "--max_subs=0", "--max_msgs=0", "--max_bytes=0", "--max_age=0s", "--hbi=10s", "--hbt=3s", "--hbf=1", "-SDV"]