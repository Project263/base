-- +goose Up
INSERT INTO users (id, username, email, password) VALUES
(gen_random_uuid(), 'admin', 'admin@example.com', 'hashed_password'),
(gen_random_uuid(), 'user1', 'user1@example.com', 'hashed_password');

-- +goose Down
DELETE FROM users WHERE email IN ('admin@example.com', 'user1@example.com');
