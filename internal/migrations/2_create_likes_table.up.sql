-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE likes (
                       id SERIAL PRIMARY KEY,
                       user_id INT NOT NULL,
                       post_id INT NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
