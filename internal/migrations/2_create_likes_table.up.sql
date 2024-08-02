-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE likes (
                       id SERIAL PRIMARY KEY,
                       liker_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                       liked_user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                       UNIQUE (liker_id, liked_user_id) -- предотвращает повторные лайки одним пользователем
);
