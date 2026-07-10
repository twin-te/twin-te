BEGIN;

CREATE TABLE auth_challenges (
    id uuid NOT NULL,
    provider text NOT NULL,
    nonce text NOT NULL,
    expired_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (provider) REFERENCES auth_providers(provider) ON UPDATE CASCADE
);

CREATE INDEX ON auth_challenges (expired_at);

COMMIT;
