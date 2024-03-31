BEGIN;

CREATE TABLE public.session (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    expired_at timestamp(3) without time zone NOT NULL
);

ALTER TABLE ONLY public.session
    ADD CONSTRAINT session_pkey PRIMARY KEY (id);

CREATE INDEX "session.id_expired_at_index" ON public.session USING btree (id, expired_at);

COMMIT;
