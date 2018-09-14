FROM alpine:latest

RUN apk --no-cache add wget ca-certificates && \
  wget https://github.com/nats-io/nats-streaming-server/releases/download/v0.11.0/nats-streaming-server-v0.11.0-linux-amd64.zip && \
  unzip nats-streaming-server-v0.11.0-linux-amd64.zip && \
  mv nats-streaming-server-v0.11.0-linux-amd64/nats-streaming-server /

EXPOSE 4222 8222

ENTRYPOINT ["/nats-streaming-server"]
CMD ["-m", "8222", "--store=FILE", "--dir=/data", "--max_channels=0", "--max_subs=0", "--max_msgs=0", "--max_bytes=0", "--max_age=0s", "--hbi=10s", "--hbt=3s", "--hbf=1", "-SDV"]