CREATE VIEW TempTable AS SELECT * FROM course_schedules WHERE module = 'Annual';

INSERT INTO course_schedules (module, day, period, room, course_id)
SELECT 'SpringA', day, period, room, course_id FROM TempTable;

INSERT INTO course_schedules (module, day, period, room, course_id)
SELECT 'SpringB', day, period, room, course_id FROM TempTable;

INSERT INTO course_schedules (module, day, period, room, course_id)
SELECT 'SpringC', day, period, room, course_id FROM TempTable;

INSERT INTO course_schedules (module, day, period, room, course_id)
SELECT 'SummerVacation', day, period, room, course_id FROM TempTable;

INSERT INTO course_schedules (module, day, period, room, course_id)
SELECT 'FallA', day, period, room, course_id FROM TempTable;

INSERT INTO course_schedules (module, day, period, room, course_id)
SELECT 'FallB', day, period, room, course_id FROM TempTable;

INSERT INTO course_schedules (module, day, period, room, course_id)
SELECT 'FallC', day, period, room, course_id FROM TempTable;

INSERT INTO course_schedules (module, day, period, room, course_id)
SELECT 'SpringVacation', day, period, room, course_id FROM TempTable;

DELETE FROM course_schedules WHERE module = 'Annual';

DROP VIEW TempTable;
