-- Table: users
CREATE TABLE users (
    user_id   SERIAL PRIMARY KEY,
    fullname  VARCHAR(255) NOT NULL,
    email     VARCHAR(255) UNIQUE NOT NULL
    -- password TEXT  -- kalau nanti mau dipakai
);

-- Table: questions
CREATE TABLE questions (
    question_id SERIAL PRIMARY KEY,
    content     TEXT NOT NULL,
    author_id   INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    created_at  TIMESTAMP DEFAULT NOW()
);

-- Table: answers
CREATE TABLE answers (
    answer_id   SERIAL PRIMARY KEY,
    content     TEXT NOT NULL,
    author_id   INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    question_id INT NOT NULL REFERENCES questions(question_id) ON DELETE CASCADE,
    created_at  TIMESTAMP DEFAULT NOW()
);
