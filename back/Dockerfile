FROM golang:1.24 AS local
WORKDIR /usr/src/twin-te/back

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
CMD [ "go", "tool", "air" ]

FROM golang:1.24-bullseye AS builder
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o /usr/local/bin/app

FROM debian:bullseye-slim AS deploy

RUN apt-get update \
  && DEBIAN_FRONTEND=noninteractive apt-get install -yq ca-certificates openssl \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder /usr/local/bin/app .

EXPOSE 8080

CMD ["./app", "serve"]
