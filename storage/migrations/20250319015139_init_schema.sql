-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    login VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    nickname VARCHAR NOT NULL,
    avatar VARCHAR,
    advanced_version BOOLEAN DEFAULT false,
    phone VARCHAR UNIQUE,
    is_verified_phone BOOLEAN DEFAULT false,
    email VARCHAR UNIQUE NOT NULL,
    is_verified_mail BOOLEAN DEFAULT false,
    age INT,
    height INT,
    weight INT,
    sex VARCHAR,
    day_streak INT DEFAULT 0,
    is_train_today BOOLEAN DEFAULT false,
    points INT DEFAULT 0,
    created_at BIGINT,
    update_at BIGINT
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS achievements (
    id SERIAL PRIMARY KEY,
    title VARCHAR UNIQUE NOT NULL,
    description VARCHAR NOT NULL,
    image VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_achievements (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    achievement_id INT REFERENCES achievements(id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS muscles (
    id SERIAL PRIMARY KEY,
    title VARCHAR UNIQUE NOT NULL,
    image VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS exercises (
    id SERIAL PRIMARY KEY,
    title VARCHAR UNIQUE NOT NULL,
    description VARCHAR UNIQUE NOT NULL,
    image VARCHAR,
    video_url VARCHAR,
    equipment_id INT NOT NULL,
    sets INT NOT NULL,
    reps INT NOT NULL,
    difficult INT NOT NULL,
    lead_muscle_id INT NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS equipments (
    id SERIAL PRIMARY KEY,
    title VARCHAR UNIQUE NOT NULL,
    image VARCHAR
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS trains (
    id SERIAL PRIMARY KEY,
    title VARCHAR UNIQUE NOT NULL,
    description VARCHAR NOT NULL,
    image VARCHAR,
    video_url VARCHAR,
    difficult INT,
    duration_time INT,
    lead_muscle_id INT NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS trains_exercises (
    id SERIAL PRIMARY KEY,
    train_id INT REFERENCES trains(id),
    exercise_id INT REFERENCES exercises(id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS exercises_help_muscle (
    id SERIAL PRIMARY KEY,
    exercise_id INT REFERENCES exercises(id),
    help_muscle_id INT REFERENCES muscles(id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS train_help_muscle (
    id SERIAL PRIMARY KEY,
    train_id INT REFERENCES trains(id),
    help_muscle_id INT REFERENCES muscles(id)
);
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
DROP TABLE IF EXISTS equipments;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS exercises;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS achievements;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS users_achievements;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS muscles;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS trains;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS trains_exercises;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS exercises_help_muscle;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS train_help_muscle;
-- +goose StatementEnd
