CREATE TABLE users (
    id          BIGSERIAL PRIMARY KEY,
    google_id   TEXT        NOT NULL,
    email       TEXT        NOT NULL,
    name        TEXT        NOT NULL,
    picture     TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at  TIMESTAMPTZ
);

CREATE UNIQUE INDEX ux_users_google_id ON users (google_id);
CREATE UNIQUE INDEX ux_users_email     ON users (email);
CREATE INDEX ix_users_deleted_at       ON users (deleted_at);
