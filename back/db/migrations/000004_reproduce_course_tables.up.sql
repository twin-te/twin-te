BEGIN;

CREATE TYPE public.course_methods_method_enum AS ENUM (
    'OnlineAsynchronous',
    'OnlineSynchronous',
    'FaceToFace',
    'Others'
);

CREATE TYPE public.course_schedules_day_enum AS ENUM (
    'Sun',
    'Mon',
    'Tue',
    'Wed',
    'Thu',
    'Fri',
    'Sat',
    'Intensive',
    'Appointment',
    'AnyTime',
    'Unknown'
);

CREATE TYPE public.course_schedules_module_enum AS ENUM (
    'SpringA',
    'SpringB',
    'SpringC',
    'FallA',
    'FallB',
    'FallC',
    'SummerVacation',
    'SpringVacation',
    'Annual',
    'Unknown'
);

CREATE TABLE public.course_methods (
    id integer NOT NULL,
    method public.course_methods_method_enum NOT NULL,
    course_id uuid
);

CREATE SEQUENCE public.course_methods_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.course_methods_id_seq OWNED BY public.course_methods.id;

CREATE TABLE public.course_recommended_grades (
    id integer NOT NULL,
    grade smallint NOT NULL,
    course_id uuid
);

CREATE SEQUENCE public.course_recommended_grades_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.course_recommended_grades_id_seq OWNED BY public.course_recommended_grades.id;

CREATE TABLE public.course_schedules (
    id integer NOT NULL,
    module public.course_schedules_module_enum NOT NULL,
    day public.course_schedules_day_enum NOT NULL,
    period smallint NOT NULL,
    room text NOT NULL,
    course_id uuid
);

CREATE SEQUENCE public.course_schedules_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.course_schedules_id_seq OWNED BY public.course_schedules.id;

CREATE TABLE public.courses (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    year smallint NOT NULL,
    code text NOT NULL,
    name text NOT NULL,
    instructor text NOT NULL,
    credit numeric NOT NULL,
    overview text NOT NULL,
    remarks text NOT NULL,
    last_update timestamp with time zone NOT NULL,
    has_parse_error boolean NOT NULL,
    is_annual boolean DEFAULT false NOT NULL
);

ALTER TABLE ONLY public.course_methods ALTER COLUMN id SET DEFAULT nextval('public.course_methods_id_seq'::regclass);

ALTER TABLE ONLY public.course_recommended_grades ALTER COLUMN id SET DEFAULT nextval('public.course_recommended_grades_id_seq'::regclass);

ALTER TABLE ONLY public.course_schedules ALTER COLUMN id SET DEFAULT nextval('public.course_schedules_id_seq'::regclass);

ALTER TABLE ONLY public.courses
    ADD CONSTRAINT "PK_3f70a487cc718ad8eda4e6d58c9" PRIMARY KEY (id);

ALTER TABLE ONLY public.course_methods
    ADD CONSTRAINT "PK_412b40a7891d105d69847f77f72" PRIMARY KEY (id);

ALTER TABLE ONLY public.course_schedules
    ADD CONSTRAINT "PK_68118fc569f0c9ebb03fb79f80e" PRIMARY KEY (id);

ALTER TABLE ONLY public.course_recommended_grades
    ADD CONSTRAINT "PK_832c34a40a52a56375e2796b3c6" PRIMARY KEY (id);

CREATE UNIQUE INDEX "IDX_68ca51dc447bc2c03d5f1c44b8" ON public.courses USING btree (year, code);

ALTER TABLE ONLY public.course_recommended_grades
    ADD CONSTRAINT "FK_54fbcd6eb4ed154ac8be7e7dc72" FOREIGN KEY (course_id) REFERENCES public.courses(id);

ALTER TABLE ONLY public.course_methods
    ADD CONSTRAINT "FK_b65806d6baf43fbfa16751d9d5c" FOREIGN KEY (course_id) REFERENCES public.courses(id);

ALTER TABLE ONLY public.course_schedules
    ADD CONSTRAINT "FK_f0cd61820798323cb06ca91105d" FOREIGN KEY (course_id) REFERENCES public.courses(id);

COMMIT;
