CREATE TABLE IF NOT EXISTS quiz (
    id SERIAL primary key,
    name TEXT NOT NULL,
    training_id INT NOT NULL,
    CONSTRAINT fk_quiz_training FOREIGN KEY (training_id) REFERENCES trainings (id)
);