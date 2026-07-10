BEGIN;

CREATE TABLE apple_credentials (
    user_id uuid NOT NULL,
    client_id text NOT NULL,
    refresh_token text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

COMMIT;
