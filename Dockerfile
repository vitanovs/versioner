FROM golang:1.13.0 AS builder
LABEL stage="versioner-builder-env"
WORKDIR /tmp/app
COPY . .
RUN make release

FROM busybox:1.31.1
WORKDIR /app
COPY --from=builder /tmp/app/bin/versioner .
ENTRYPOINT [ "./versioner" ]
