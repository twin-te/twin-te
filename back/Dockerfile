FROM golang:1.23 AS local
WORKDIR /usr/src/twin-te/back

RUN go install golang.org/x/tools/gopls@v0.18.1 \
  && go install github.com/bufbuild/buf/cmd/buf@latest \
  && go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest \
  && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
  && go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest \
  && go install github.com/spf13/cobra-cli@latest \
  && go install github.com/air-verse/air@v1.61.7 \
  && go install github.com/go-delve/delve/cmd/dlv@latest

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
CMD [ "air" ]

FROM golang:1.23-bullseye AS builder
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
