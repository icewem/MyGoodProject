-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) NOT NULL UNIQUE,
                       likes_count INTEGER DEFAULT 0
);
