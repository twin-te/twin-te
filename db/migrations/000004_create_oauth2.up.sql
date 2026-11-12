BEGIN;

CREATE TABLE oauth2_clients (
    id text NOT NULL,
    name text NOT NULL,
    client_uri text NOT NULL,
    redirect_uris jsonb NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE oauth2_requests (
    kind text NOT NULL,
    signature text NOT NULL,
    request_id text NOT NULL,
    user_id uuid NOT NULL,
    client_id text NOT NULL,
    requested_at timestamp without time zone NOT NULL,
    requested_scopes jsonb,
    granted_scopes jsonb,
    form jsonb,
    session jsonb NOT NULL,
    active boolean NOT NULL DEFAULT true,
    expires_at timestamp without time zone NOT NULL,
    PRIMARY KEY (kind, signature),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (client_id) REFERENCES oauth2_clients(id) ON DELETE CASCADE
);

CREATE INDEX oauth2_requests_request_id_idx ON oauth2_requests (request_id);
CREATE INDEX oauth2_requests_user_id_client_id_idx ON oauth2_requests (user_id, client_id);
CREATE INDEX oauth2_requests_expires_at_idx ON oauth2_requests (expires_at);

COMMIT;
