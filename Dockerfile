FROM golang:1.22.9 AS builder

WORKDIR /workspace
COPY . .

RUN CGO_ENABLED=0 GOOS=linux make build

FROM alpine:3.18

WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=builder /workspace/build/* ./
COPY --from=builder /workspace/config ./config
COPY --from=builder /workspace/pkg/migration ./pkg/migration


