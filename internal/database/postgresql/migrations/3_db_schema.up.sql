BEGIN;


UPDATE schema_migrations SET dirty=false;

CREATE TABLE links (
    id              BIGSERIAL PRIMARY KEY,
    original_link   TEXT NOT NULL UNIQUE,
    short_link      TEXT NOT NULL UNIQUE,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE INDEX idx_short_link ON links (short_link);

COMMIT;