BEGIN;

CREATE TABLE already_reads (
    id uuid NOT NULL,
    user_id uuid NOT NULL,    
    announcement_id uuid NOT NULL,
    read_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (user_id, announcement_id)
);

CREATE TABLE auth_providers (
    provider text NOT NULL,
    PRIMARY KEY (provider)
);

INSERT INTO auth_providers VALUES
    ('Google'),
    ('Apple'),
    ('Twitter');

CREATE TABLE users (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone,
    updated_at timestamp without time zone NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE user_authentications (
    user_id uuid NOT NULL,
    provider text NOT NULL,
    social_id text NOT NULL,
    UNIQUE (provider, social_id),
    PRIMARY KEY (user_id, provider, social_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (provider) REFERENCES auth_providers(provider) ON UPDATE CASCADE
);

CREATE TABLE sessions (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    expired_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE payment_users (
    id text NOT NULL,
    user_id uuid NOT NULL,
    display_name text,
    link text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,    
    PRIMARY KEY (id)
);

CREATE INDEX ON payment_users (user_id);

CREATE TABLE timetable_modules (
    module text NOT NULL,
    PRIMARY KEY (module)
);

INSERT INTO timetable_modules VALUES
    ('SpringA'),
    ('SpringB'),
    ('SpringC'),
    ('FallA'),
    ('FallB'),
    ('FallC'),
    ('SummerVacation'),
    ('SpringVacation');

CREATE TABLE timetable_days (
    day text NOT NULL,
    PRIMARY KEY (day)
);

INSERT INTO timetable_days VALUES
    ('Sun'),
    ('Mon'),
    ('Tue'),
    ('Wed'),
    ('Thu'),
    ('Fri'),
    ('Sat'),
    ('Intensive'),
    ('Appointment'),
    ('AnyTime'),
    ('NT');    

CREATE TABLE timetable_methods (
    method text NOT NULL,
    PRIMARY KEY (method)
);

INSERT INTO timetable_methods VALUES
    ('OnlineAsynchronous'),
    ('OnlineSynchronous'),
    ('FaceToFace'),
    ('Others');

CREATE TABLE courses (
    id uuid NOT NULL,
    year smallint NOT NULL,
    code text NOT NULL,
    name text NOT NULL,
    instructors text NOT NULL,
    credit numeric(4, 1) NOT NULL,
    overview text NOT NULL,
    remarks text NOT NULL,
    last_updated_at timestamp without time zone NOT NULL,
    has_parse_error boolean NOT NULL,
    is_annual boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,    
    PRIMARY KEY (id),
    UNIQUE (year, code)
);

CREATE TABLE course_methods (
    course_id uuid NOT NULL,    
    method text NOT NULL,
    PRIMARY KEY (course_id, method),
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE,
    FOREIGN KEY (method) REFERENCES timetable_methods(method) ON UPDATE CASCADE
);

CREATE TABLE course_recommended_grades (
    course_id uuid NOT NULL,
    recommended_grade smallint NOT NULL,
    PRIMARY KEY (course_id, recommended_grade),
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
);

CREATE TABLE course_schedules (
    course_id uuid NOT NULL,
    module text NOT NULL,
    day text NOT NULL,
    period smallint NOT NULL,
    locations text NOT NULL,
    PRIMARY KEY (course_id, module, day, period, locations),
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE,
    FOREIGN KEY (module) REFERENCES timetable_modules(module) ON UPDATE CASCADE,
    FOREIGN KEY (day) REFERENCES timetable_days(day) ON UPDATE CASCADE
);

CREATE TABLE registered_courses (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    year smallint NOT NULL,
    course_id uuid,
    name text,
    instructors text,
    credit numeric(4, 1),
    methods text[],
    schedules jsonb,
    memo text NOT NULL,
    attendance smallint NOT NULL,
    absence smallint NOT NULL,
    late smallint NOT NULL,
    UNIQUE (user_id, course_id),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,    
    PRIMARY KEY (id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
);

CREATE INDEX ON registered_courses (user_id, year);

CREATE TABLE registered_course_tags (
    registered_course_id uuid NOT NULL,
    tag_id uuid NOT NULL,
    PRIMARY KEY (registered_course_id, tag_id),
    FOREIGN KEY (registered_course_id) REFERENCES registered_courses(id) ON DELETE CASCADE
);

CREATE TABLE tags (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    name text NOT NULL,
    "order" smallint NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,    
    UNIQUE (user_id, "order") DEFERRABLE INITIALLY DEFERRED,
    PRIMARY KEY (id)
);

COMMIT;
