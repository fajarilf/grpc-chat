CREATE TABLE IF NOT EXISTS rooms (
    id              SERIAL          PRIMARY KEY,
    name            TEXT            NOT NULL,
    description     TEXT            NOT NULL DEFAULT '',
    background      TEXT            NOT NULL DEFAULT '',
    -- background_type: 0 = COLOR, 1 = IMAGE (mirrors the BgTypes proto enum).
    background_type SMALLINT        NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT now()
);
