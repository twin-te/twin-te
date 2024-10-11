BEGIN;

DROP TABLE tags CASCADE;

DROP TABLE registered_course_tag_ids CASCADE;
DROP TABLE registered_courses CASCADE;

DROP TABLE course_schedules CASCADE;
DROP TABLE course_recommended_grades CASCADE;
DROP TABLE course_methods CASCADE;
DROP TABLE courses CASCADE;

DROP TABLE timetable_methods CASCADE;
DROP TABLE timetable_days CASCADE;
DROP TABLE timetable_modules CASCADE;

DROP TABLE payment_users CASCADE;

DROP TABLE sessions CASCADE;

DROP TABLE user_authentications CASCADE;
DROP TABLE users CASCADE;

DROP TABLE auth_providers CASCADE;

DROP TABLE already_reads CASCADE;

COMMIT;
