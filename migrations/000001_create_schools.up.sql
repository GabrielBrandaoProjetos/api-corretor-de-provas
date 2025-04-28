CREATE TABLE IF NOT EXISTS schools (
    id          text            PRIMARY KEY,
    name        text            NOT NULL,
    register    text            NOT NULL UNIQUE,
    unit        text            NOT NULL,
    address     text            NOT NULL,
    created_at  timestamp       NOT NULL DEFAULT NOW(),
    updated_at  timestamp       NOT NULL DEFAULT NOW()
);