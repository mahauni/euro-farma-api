CREATE TABLE IF NOT EXISTS trainings (
    id SERIAL primary key,
    name TEXT NOT NULL,
    category_id INT NOT NULL,
    CONSTRAINT fk_training_category FOREIGN KEY (category_id) REFERENCES categories (id)
);