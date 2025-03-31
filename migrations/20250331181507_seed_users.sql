-- +goose Up
INSERT INTO users (id, username, email, password, role) VALUES
(gen_random_uuid(), 'admin', 'admin@example.com', 'admin', 'admin'),
(gen_random_uuid(), 'user1', 'user1@example.com', '123', 'user');

-- +goose Down
DELETE FROM users WHERE email IN ('admin@example.com', 'user1@example.com');
