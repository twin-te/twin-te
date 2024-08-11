FROM golang:1 AS db-migration
WORKDIR /app

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD [ "sh", "-c", "migrate -database $DB_URL -path ./migrations up" ]
