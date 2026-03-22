FROM golang:1.22-alpine AS builder

ARG VERSION=dev
WORKDIR /build

COPY go.mod go.sum* ./
RUN go mod download || true

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -trimpath \
    -ldflags "-w -s -extldflags '-static' -X 'main.Version=${VERSION}'" \
    -o d2-live .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/d2-live /d2-live

USER 65534:65534

EXPOSE 8090

ENTRYPOINT ["/d2-live"]
