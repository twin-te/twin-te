BEGIN;

DROP TABLE IF EXISTS public.course_schedules;

DROP TABLE IF EXISTS public.course_recommended_grades;

DROP TABLE IF EXISTS public.course_methods;

DROP INDEX IF EXISTS public."IDX_68ca51dc447bc2c03d5f1c44b8";

DROP TABLE IF EXISTS public.courses;

DROP SEQUENCE IF EXISTS public.course_schedules_id_seq;

DROP SEQUENCE IF EXISTS public.course_recommended_grades_id_seq;

DROP SEQUENCE IF EXISTS public.course_methods_id_seq;

DROP TYPE IF EXISTS public.course_schedules_module_enum;

DROP TYPE IF EXISTS public.course_schedules_day_enum;

DROP TYPE IF EXISTS public.course_methods_method_enum;

COMMIT;
