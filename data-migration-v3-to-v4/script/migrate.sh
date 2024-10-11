#!/usr/bin/env bash

POSTGRES_URL="${POSTGRES_URL:-postgres://postgres:password@db:5432/twinte_db?sslmode=disable}"

psql -d $POSTGRES_URL -c "\COPY users FROM '/tmp/data/users.csv' WITH (FORMAT csv, HEADER, QUOTE '\"', ESCAPE '\', NULL 'null')"
psql -d $POSTGRES_URL -c "\COPY user_authentications FROM '/tmp/data/user_authentications.csv' WITH (FORMAT csv, HEADER, QUOTE '\"', ESCAPE '\', NULL 'null')"

psql -d $POSTGRES_URL -c "\COPY courses FROM '/tmp/data/courses_found.csv' WITH (FORMAT csv, HEADER, QUOTE '\"', ESCAPE '\', NULL 'null')"
psql -d $POSTGRES_URL -c "\COPY courses FROM '/tmp/data/courses_not_found.csv' WITH (FORMAT csv, HEADER, QUOTE '\"', ESCAPE '\', NULL 'null')"
psql -d $POSTGRES_URL -c "\COPY course_methods FROM '/tmp/data/course_methods_found.csv' WITH (FORMAT csv, HEADER, QUOTE '\"', ESCAPE '\', NULL 'null')"
psql -d $POSTGRES_URL -c "\COPY course_methods FROM '/tmp/data/course_methods_not_found.csv' WITH (FORMAT csv, HEADER, QUOTE '\"', ESCAPE '\', NULL 'null')"
psql -d $POSTGRES_URL -c "\COPY course_recommended_grades FROM '/tmp/data/course_recommended_grades_found.csv' WITH (FORMAT csv, HEADER, QUOTE '\"', ESCAPE '\', NULL 'null')"
psql -d $POSTGRES_URL -c "\COPY course_recommended_grades FROM '/tmp/data/course_recommended_grades_not_found.csv' WITH (FORMAT csv, HEADER, QUOTE '\"', ESCAPE '\', NULL 'null')"
psql -d $POSTGRES_URL -c "\COPY course_schedules FROM '/tmp/data/course_schedules_found.csv' WITH (FORMAT csv, HEADER, QUOTE '\"', ESCAPE '\', NULL 'null')"
psql -d $POSTGRES_URL -c "\COPY course_schedules FROM '/tmp/data/course_schedules_not_found.csv' WITH (FORMAT csv, HEADER, QUOTE '\"', ESCAPE '\', NULL 'null')"

psql -d $POSTGRES_URL -c "\COPY tags FROM '/tmp/data/tags.csv' WITH (FORMAT csv, HEADER, QUOTE '\"', ESCAPE '\', NULL 'null')"

psql -d $POSTGRES_URL -c "\COPY registered_courses FROM '/tmp/data/registered_courses.csv' WITH (FORMAT csv, HEADER, QUOTE '\"', ESCAPE '\', NULL 'null')"
psql -d $POSTGRES_URL -c "\COPY registered_course_tag_ids FROM '/tmp/data/registered_course_tag_ids.csv' WITH (FORMAT csv, HEADER, QUOTE '\"', ESCAPE '\', NULL 'null')"

psql -d $POSTGRES_URL -c "\COPY payment_users FROM '/tmp/data/payment_users.csv' WITH (FORMAT csv, HEADER, QUOTE '\"', ESCAPE '\', NULL 'null')"

psql -d $POSTGRES_URL -c "\COPY sessions FROM '/tmp/data/sessions.csv' WITH (FORMAT csv, HEADER, QUOTE '\"', ESCAPE '\', NULL 'null')"
