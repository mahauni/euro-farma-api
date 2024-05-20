CREATE TABLE IF NOT EXISTS quiz_questions (
    id SERIAL primary key,
    question TEXT NOT NULL,
    score INT NOT NULl,
    quiz_id INT NOT NULL,
    CONSTRAINT fk_question_quiz FOREIGN KEY (quiz_id) REFERENCES quiz (id)
);