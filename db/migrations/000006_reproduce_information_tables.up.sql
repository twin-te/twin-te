BEGIN;

CREATE TABLE public.already_reads (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    information_id uuid NOT NULL,
    read_user text NOT NULL,
    read_at timestamp without time zone NOT NULL
);

CREATE TABLE public.information (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    title character varying NOT NULL,
    content text NOT NULL,
    published_at timestamp without time zone NOT NULL
);

ALTER TABLE ONLY public.information
    ADD CONSTRAINT "PK_091c910b61c3170a50eaf22e0c4" PRIMARY KEY (id);

ALTER TABLE ONLY public.already_reads
    ADD CONSTRAINT "PK_7b7e4dc4f3dac9076561acd7832" PRIMARY KEY (id);

COMMIT;
