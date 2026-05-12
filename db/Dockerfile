FROM golang:1@sha256:2981696eed011d747340d7252620932677929cce7d2d539602f56a8d7e9b660b AS migration

WORKDIR /usr/src/twin-te/db

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY ./migrations ./
