
CREATE TABLE IF NOT EXISTS users (
    id SERIAL primary key,
    name TEXT NOT NULL,
    password TEXT NOT NULl,
    email TEXT NOT NULl
);