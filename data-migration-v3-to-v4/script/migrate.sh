#!/usr/bin/env bash

POSTGRES_URL="postgres://postgres:password@db:5432/twinte_db?sslmode=disable"

psql -d $POSTGRES_URL -c "\COPY courses FROM '/tmp/courses_found.csv' WITH CSV NULL 'null';"
psql -d $POSTGRES_URL -c "\COPY courses FROM '/tmp/courses_not_found.csv' WITH CSV NULL 'null';"
psql -d $POSTGRES_URL -c "\COPY course_methods FROM '/tmp/course_methods_found.csv' WITH CSV NULL 'null';"
psql -d $POSTGRES_URL -c "\COPY course_methods FROM '/tmp/course_methods_not_found.csv' WITH CSV NULL 'null';"
psql -d $POSTGRES_URL -c "\COPY course_recommended_grades FROM '/tmp/course_recommended_grades_found.csv' WITH CSV NULL 'null';"
psql -d $POSTGRES_URL -c "\COPY course_recommended_grades FROM '/tmp/course_recommended_grades_not_found.csv' WITH CSV NULL 'null';"
psql -d $POSTGRES_URL -c "\COPY course_schedules FROM '/tmp/course_schedules_found.csv' WITH CSV NULL 'null';"
psql -d $POSTGRES_URL -c "\COPY course_schedules FROM '/tmp/course_schedules_not_found.csv' WITH CSV NULL 'null';"
psql -d $POSTGRES_URL -c "\COPY payment_users FROM '/tmp/payment_users.csv' WITH CSV NULL 'null';"
psql -d $POSTGRES_URL -c "\COPY registered_courses FROM '/tmp/registered_courses.csv' WITH CSV NULL 'null';"
psql -d $POSTGRES_URL -c "\COPY registered_course_tags FROM '/tmp/registered_course_tags.csv' WITH CSV NULL 'null';"
psql -d $POSTGRES_URL -c "\COPY sessions FROM '/tmp/sessions.csv' WITH CSV NULL 'null';"
psql -d $POSTGRES_URL -c "\COPY tags FROM '/tmp/tags.csv' WITH CSV NULL 'null';"
psql -d $POSTGRES_URL -c "\COPY users FROM '/tmp/users.csv' WITH CSV NULL 'null';"
psql -d $POSTGRES_URL -c "\COPY user_authentications FROM '/tmp/user_authentications.csv' WITH CSV NULL 'null';"
