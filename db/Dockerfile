FROM golang:1 as migration

WORKDIR /usr/src/twin-te/db

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY ./migrations ./
