BEGIN;

CREATE TABLE public.payment_users (
    id text NOT NULL,
    twinte_user_id uuid NOT NULL,
    display_name text,
    link text
);

ALTER TABLE ONLY public.payment_users
    ADD CONSTRAINT "PK_1438bd4715f036ae4353ae95505" PRIMARY KEY (id);

COMMIT;
