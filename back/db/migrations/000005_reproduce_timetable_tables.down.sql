BEGIN;

DROP INDEX IF EXISTS public."IDX_c102260e1ffed113fef8ead148";

DROP INDEX IF EXISTS public."IDX_41594bc21c7a7adff5f2a4574c";

DROP TABLE IF EXISTS public.registered_course_tags;

DROP TABLE IF EXISTS public.tags;

DROP INDEX IF EXISTS public."IDX_fbc9587b218000acc37d7c6385";

DROP TABLE IF EXISTS public.registered_courses;

DROP TYPE IF EXISTS public.registered_courses_methods_enum;

COMMIT;