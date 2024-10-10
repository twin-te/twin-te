#!/usr/bin/env bash

cd $(dirname $0) || exit
cd ../

COPY courses FROM '/tmp/courses_found.csv' WITH CSV NULL 'null';
COPY courses FROM '/tmp/courses_not_found.csv' WITH CSV NULL 'null';
COPY course_methods FROM '/tmp/course_methods_found.csv' WITH CSV NULL 'null';
COPY course_methods FROM '/tmp/course_methods_not_found.csv' WITH CSV NULL 'null';
COPY course_recommended_grades FROM '/tmp/course_recommended_grades_found.csv' WITH CSV NULL 'null';
COPY course_recommended_grades FROM '/tmp/course_recommended_grades_not_found.csv' WITH CSV NULL 'null';
COPY course_schedules FROM '/tmp/course_schedules_found.csv' WITH CSV NULL 'null';
COPY course_schedules FROM '/tmp/course_schedules_not_found.csv' WITH CSV NULL 'null';
COPY payment_users FROM '/tmp/payment_users.csv' WITH CSV NULL 'null';
COPY registered_courses FROM '/tmp/registered_courses.csv' WITH CSV NULL 'null';
COPY registered_course_tags FROM '/tmp/registered_course_tags.csv' WITH CSV NULL 'null';
COPY sessions FROM '/tmp/sessions.csv' WITH CSV NULL 'null';
COPY tags FROM '/tmp/tags.csv' WITH CSV NULL 'null';
COPY users FROM '/tmp/users.csv' WITH CSV NULL 'null';
COPY user_authentications FROM '/tmp/user_authentications.csv' WITH CSV NULL 'null';
