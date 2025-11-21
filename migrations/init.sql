CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    pass_hash TEXT NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO users (username, pass_hash, is_admin)
VALUES ('DonaMaria', 'SeuCantinho123', TRUE)
ON CONFLICT (username) DO NOTHING;
