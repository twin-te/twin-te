ARG GO_VERSION=1.26.2
FROM golang:${GO_VERSION}@sha256:b54cbf583d390341599d7bcbc062425c081105cc5ef6d170ced98ef9d047c716 AS base
WORKDIR /usr/src/twin-te/back

RUN --mount=type=cache,target=/go/pkg/mod/,sharing=locked \
  --mount=type=bind,source=go.sum,target=go.sum \
  --mount=type=bind,source=go.mod,target=go.mod \
  go mod download -x

FROM base AS local

CMD [ "go", "tool", "air" ]

FROM golang:${GO_VERSION}-trixie@sha256:4a7137ea573f79c86ae451ff05817ed762ef5597fcf732259e97abeb3108d873 AS builder
WORKDIR /usr/src/app

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,target=. \
  GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o /usr/local/bin/app

FROM debian:trixie-slim@sha256:109e2c65005bf160609e4ba6acf7783752f8502ad218e298253428690b9eaa4b AS deploy

RUN --mount=type=cache,target=/var/lib/apt,sharing=locked \
  --mount=type=cache,target=/var/cache/apt,sharing=locked \
  apt-get update && apt-get install -y ca-certificates openssl

COPY --from=builder /usr/local/bin/app .

EXPOSE 8080

CMD ["./app", "serve"]
