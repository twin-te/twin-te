BEGIN;

DROP INDEX IF EXISTS public."session.id_expired_at_index";

DROP TABLE IF EXISTS public.session;

COMMIT;