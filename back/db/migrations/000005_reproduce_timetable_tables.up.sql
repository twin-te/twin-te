BEGIN;

CREATE TYPE public.registered_courses_methods_enum AS ENUM (
    'OnlineAsynchronous',
    'OnlineSynchronous',
    'FaceToFace',
    'Others'
);

CREATE TABLE public.registered_course_tags (
    tag uuid NOT NULL,
    registered_course uuid NOT NULL
);

CREATE TABLE public.registered_courses (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    year smallint NOT NULL,
    course_id uuid,
    name text,
    instractor text,
    credit numeric,
    methods public.registered_courses_methods_enum[],
    schedules jsonb,
    memo text NOT NULL,
    attendance integer NOT NULL,
    absence integer NOT NULL,
    late integer NOT NULL
);

CREATE TABLE public.tags (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    name text NOT NULL,
    "position" integer NOT NULL
);

ALTER TABLE ONLY public.registered_courses
    ADD CONSTRAINT "PK_1aa4b72144999c7ac6f37e52fca" PRIMARY KEY (id);

ALTER TABLE ONLY public.registered_course_tags
    ADD CONSTRAINT "PK_605d76e7c4a818b543daae4cdf2" PRIMARY KEY (tag, registered_course);

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT "PK_e7dc17249a1148a1970748eda99" PRIMARY KEY (id);

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT "UQ_b861cf6ec9af09190780481a311" UNIQUE (user_id, "position") DEFERRABLE INITIALLY DEFERRED;

CREATE INDEX "IDX_41594bc21c7a7adff5f2a4574c" ON public.registered_course_tags USING btree (registered_course);

CREATE INDEX "IDX_c102260e1ffed113fef8ead148" ON public.registered_course_tags USING btree (tag);

CREATE UNIQUE INDEX "IDX_fbc9587b218000acc37d7c6385" ON public.registered_courses USING btree (user_id, course_id);

ALTER TABLE ONLY public.registered_course_tags
    ADD CONSTRAINT "FK_41594bc21c7a7adff5f2a4574cc" FOREIGN KEY (registered_course) REFERENCES public.registered_courses(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.registered_course_tags
    ADD CONSTRAINT "FK_c102260e1ffed113fef8ead1481" FOREIGN KEY (tag) REFERENCES public.tags(id) ON UPDATE CASCADE ON DELETE CASCADE;

COMMIT;
