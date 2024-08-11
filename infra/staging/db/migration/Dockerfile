FROM migrate/migrate

RUN apk add --no-cache git

COPY migration.sh /usr/local/bin
RUN chmod +x /usr/local/bin/migration.sh

ENTRYPOINT [ "migration.sh" ]
