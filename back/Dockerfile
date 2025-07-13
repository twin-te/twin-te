ARG GO_VERSION=1.24.5
FROM golang:${GO_VERSION} AS base
WORKDIR /usr/src/twin-te/back

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,source=go.mod,target=go.mod \
  go mod download -x

FROM base AS local

CMD [ "go", "tool", "air" ]

FROM golang:${GO_VERSION}-bullseye AS builder
WORKDIR /usr/src/app

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,target=. \
  GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o /usr/local/bin/app

FROM debian:bullseye-slim AS deploy

RUN --mount=type=cache,target=/var/lib/apt,sharing=locked \
  --mount=type=cache,target=/var/cache/apt,sharing=locked \
  apt-get update && apt-get install -y ca-certificates openssl

COPY --from=builder /usr/local/bin/app .

EXPOSE 8080

CMD ["./app", "serve"]
