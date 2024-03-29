BEGIN;

CREATE TYPE public.user_authentications_provider_enum AS ENUM (
    'Google',
    'Twitter',
    'Apple'
);

CREATE TABLE public.user_authentications (
    id integer NOT NULL,
    provider public.user_authentications_provider_enum NOT NULL,
    social_id character varying NOT NULL,
    user_id uuid
);

CREATE SEQUENCE public.user_authentications_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.user_authentications_id_seq OWNED BY public.user_authentications.id;

CREATE TABLE public.users (
    id uuid NOT NULL,
    "createdAt" timestamp without time zone DEFAULT now() NOT NULL,
    "deletedAt" timestamp without time zone
);

ALTER TABLE ONLY public.user_authentications ALTER COLUMN id SET DEFAULT nextval('public.user_authentications_id_seq'::regclass);

ALTER TABLE ONLY public.user_authentications
    ADD CONSTRAINT "PK_5357fb1162b50b926c77290c8bc" PRIMARY KEY (id);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT "PK_a3ffb1c0c8416b9fc6f907b7433" PRIMARY KEY (id);

ALTER TABLE ONLY public.user_authentications
    ADD CONSTRAINT "FK_163ff5c9a502621798f57606e80" FOREIGN KEY (user_id) REFERENCES public.users(id);

COMMIT;
