CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    pass_hash TEXT NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS spaces (
    id SERIAL PRIMARY KEY,
    location TEXT NOT NULL,
    substation TEXT NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    capacity INTEGER NOT NULL,
    image BYTEA
);

INSERT INTO users (username, pass_hash, is_admin)
VALUES ('DonaMaria', '$2a$10$lAxZ5mptOG9SyHr.5KuzpuNWl6MzdvOgxhUD.GJG9aZ4cQwvEC9qC', TRUE)
ON CONFLICT (username) DO NOTHING;