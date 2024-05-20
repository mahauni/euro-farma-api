CREATE TABLE IF NOT EXISTS users (
    id SERIAL primary key,
    name TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL,
    points INT NOT NULL
);