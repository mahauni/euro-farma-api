CREATE TABLE IF NOT EXISTS training_user (
    id SERIAL primary key,
    training_id INT NOT NULL,
    user_id INT NOT NULL,
    CONSTRAINT fk_training_user_training FOREIGN KEY (training_id) REFERENCES trainings (id),
    CONSTRAINT fk_training_user_user FOREIGN KEY (user_id) REFERENCES users (id)
);