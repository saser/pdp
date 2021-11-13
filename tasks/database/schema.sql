BEGIN;

CREATE TABLE IF NOT EXISTS Tasks (
    id UUID NOT NULL,
    title TEXT NOT NULL,
    completed BOOLEAN NOT NULL,

    PRIMARY KEY (id),
    CONSTRAINT title_not_empty CHECK (title <> '')
);

COMMIT;