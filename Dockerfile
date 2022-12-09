FROM golang:1.18.0 AS builder
WORKDIR /tmp/app
COPY . .
RUN make release

FROM busybox:1.31.1
WORKDIR /app
COPY --from=builder /tmp/app/bin/versioner .
