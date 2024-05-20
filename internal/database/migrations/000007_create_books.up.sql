CREATE TABLE IF NOT EXISTS books (
    id SERIAL primary key,
    file_name TEXT NOT NULL,
    file_path TEXT NOT NULL,
    training_id INT NOT NULL,
    CONSTRAINT fk_books_training FOREIGN KEY (training_id) REFERENCES trainings (id)
);