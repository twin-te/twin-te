BEGIN;

ALTER TYPE public.course_schedules_day_enum ADD VALUE IF NOT EXISTS 'NT';

COMMIT;
