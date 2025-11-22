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
VALUES ('DonaMaria', '$2a$10$Kpbi/0XjbrAcD0C5bxcM.OO4hISNQWqAHA3pYSD10ypvJhKyEYzYW', TRUE)
ON CONFLICT (username) DO NOTHING;