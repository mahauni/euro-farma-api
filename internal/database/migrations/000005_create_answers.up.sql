CREATE TABLE IF NOT EXISTS quiz_answers (
    id SERIAL primary key,
    answer TEXT NOT NULL,
    correct BOOLEAN NOT NULl,
    quiz_id INT NOT NULL,
    question_id INT NOT NULL,
    CONSTRAINT fk_answer_quiz FOREIGN KEY (quiz_id) REFERENCES quiz (id),
    CONSTRAINT fk_answer_question FOREIGN KEY (question_id) REFERENCES quiz_questions (id)
);