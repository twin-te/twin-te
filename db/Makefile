# e.g. make migrate-create name=foo
migrate-create:
	migrate create -dir ./migrations -ext sql -seq -digits 6 ${name}

# e.g. make migrate-up db_url=${DB_URL}
migrate-up:
	migrate -database ${db_url} -path ./migrations up

# e.g. make migrate-down db_url=${DB_URL}
migrate-down:
	migrate -database ${db_url} -path ./migrations down -all

# e.g. make migrate-force db_url=${DB_URL} version=1
migrate-force:
	migrate -database ${db_url} -path ./migrations force ${version}
