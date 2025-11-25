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

CREATE TABLE IF NOT EXISTS bookings (
    id SERIAL PRIMARY KEY,
    spaceId INTEGER NOT NULL REFERENCES spaces(id),
    userId INTEGER NOT NULL REFERENCES users(id),
    bookingStart BIGINT NOT NULL,
    bookingEnd BIGINT NOT NULL
);

INSERT INTO users (username, pass_hash, is_admin)
VALUES ('DonaMaria', '$2a$10$Kpbi/0XjbrAcD0C5bxcM.OO4hISNQWqAHA3pYSD10ypvJhKyEYzYW', TRUE)
ON CONFLICT (username) DO NOTHING;


INSERT INTO spaces (id, location, substation, price, capacity, image)
VALUES (
    0,
    'Teste',
    'Substation A',
    99.99,
    5,
    NULL
);

