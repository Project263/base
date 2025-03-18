package storage

import (
	"database/sql"
	"fmt"
	"log"
)

// Выполнение SQL-запроса с проверкой ошибки
func execQuery(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	return nil
}

func InitTables(db *sql.DB) {
	// Создание всех таблиц
	if err := createUsersTable(db); err != nil {
		log.Fatalf("Ошибка создания таблицы users: %v", err)
	}
	if err := createAchievementsTable(db); err != nil {
		log.Fatalf("Ошибка создания таблицы achievements: %v", err)
	}
	if err := createUsersAchievementsTable(db); err != nil {
		log.Fatalf("Ошибка создания таблицы users_achievements: %v", err)
	}
	if err := createMusclesTable(db); err != nil {
		log.Fatalf("Ошибка создания таблицы muscles: %v", err)
	}
	if err := createExercisesTable(db); err != nil {
		log.Fatalf("Ошибка создания таблицы exercises: %v", err)
	}
	if err := createEquipmentsTable(db); err != nil {
		log.Fatalf("Ошибка создания таблицы equipments: %v", err)
	}
	if err := createTrainsTable(db); err != nil {
		log.Fatalf("Ошибка создания таблицы trains: %v", err)
	}
	if err := createTrainsExercisesTable(db); err != nil {
		log.Fatalf("Ошибка создания таблицы trains_exercises: %v", err)
	}
	if err := createExercisesHelpMuscleTable(db); err != nil {
		log.Fatalf("Ошибка создания таблицы exercises_help_muscle: %v", err)
	}
	if err := createTrainHelpMuscleTable(db); err != nil {
		log.Fatalf("Ошибка создания таблицы train_help_muscle: %v", err)
	}

	// Создание всех связей между таблицами
	if err := createExercisesRelations(db); err != nil {
		log.Fatalf("Ошибка создания связей для таблицы exercises: %v", err)
	}
	if err := createTrainsExercisesRelations(db); err != nil {
		log.Fatalf("Ошибка создания связей для таблицы trains_exercises: %v", err)
	}
	if err := createUsersAchievementsRelations(db); err != nil {
		log.Fatalf("Ошибка создания связей для таблицы users_achievements: %v", err)
	}
	if err := createTrainsRelations(db); err != nil {
		log.Fatalf("Ошибка создания связей для таблицы trains: %v", err)
	}
	if err := createExercisesHelpMuscleRelations(db); err != nil {
		log.Fatalf("Ошибка создания связей для таблицы exercises_help_muscle: %v", err)
	}
	if err := createTrainHelpMuscleRelations(db); err != nil {
		log.Fatalf("Ошибка создания связей для таблицы train_help_muscle: %v", err)
	}

	fmt.Println("Все таблицы и связи успешно созданы!")
}

// Таблица пользователей
func createUsersTable(db *sql.DB) error {
	query := `
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
    );`
	return execQuery(db, query)
}

// Таблица достижений
func createAchievementsTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS achievements (
        id SERIAL PRIMARY KEY,
        title VARCHAR UNIQUE NOT NULL,
        description VARCHAR UNIQUE NOT NULL,
        image VARCHAR NOT NULL
    );`
	return execQuery(db, query)
}

// Таблица достижений пользователей
func createUsersAchievementsTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS users_achievements (
        id SERIAL PRIMARY KEY,
        user_id INT REFERENCES users(id),
        achievement_id INT REFERENCES achievements(id)
    );`
	return execQuery(db, query)
}

// Таблица мышц
func createMusclesTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS muscles (
        id SERIAL PRIMARY KEY,
        title VARCHAR UNIQUE NOT NULL,
        image VARCHAR UNIQUE NOT NULL
    );`
	return execQuery(db, query)
}

// Таблица упражнений
func createExercisesTable(db *sql.DB) error {
	query := `
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
    );`
	return execQuery(db, query)
}

// Таблица оборудования
func createEquipmentsTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS equipments (
        id SERIAL PRIMARY KEY,
        title VARCHAR UNIQUE NOT NULL,
        image VARCHAR UNIQUE NOT NULL
    );`
	return execQuery(db, query)
}

// Таблица тренировок
func createTrainsTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS trains (
        id SERIAL PRIMARY KEY,
        title VARCHAR UNIQUE NOT NULL,
        description VARCHAR NOT NULL,
        image VARCHAR,
        video_url VARCHAR,
        difficult INT,
        duration_train INT,
        lead_muscle_id INT NOT NULL
    );`
	return execQuery(db, query)
}

// Таблица упражнений в тренировках
func createTrainsExercisesTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS trains_exercises (
        id SERIAL PRIMARY KEY,
        train_id INT REFERENCES trains(id),
        exercise_id INT REFERENCES exercises(id)
    );`
	return execQuery(db, query)
}

// Таблица вспомогательных мышц упражнений
func createExercisesHelpMuscleTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS exercises_help_muscle (
        id SERIAL PRIMARY KEY,
        exercise_id INT REFERENCES exercises(id),
        help_muscle_id INT REFERENCES muscles(id)
    );`
	return execQuery(db, query)
}

// Таблица вспомогательных мышц тренировок
func createTrainHelpMuscleTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS train_help_muscle (
        id SERIAL PRIMARY KEY,
        train_id INT REFERENCES trains(id),
        help_muscle_id INT REFERENCES muscles(id)
    );`
	return execQuery(db, query)
}

// Создание связей для таблицы exercises
func createExercisesRelations(db *sql.DB) error {
	query := `
    ALTER TABLE exercises 
        ADD FOREIGN KEY (lead_muscle_id) REFERENCES muscles(id),
        ADD FOREIGN KEY (equipment_id) REFERENCES equipments(id);
    `
	return execQuery(db, query)
}

// Создание связей для таблицы trains_exercises
func createTrainsExercisesRelations(db *sql.DB) error {
	query := `
    ALTER TABLE trains_exercises 
        ADD FOREIGN KEY (train_id) REFERENCES trains(id),
        ADD FOREIGN KEY (exercise_id) REFERENCES exercises(id);
    `
	return execQuery(db, query)
}

// Создание связей для таблицы users_achievements
func createUsersAchievementsRelations(db *sql.DB) error {
	query := `
    ALTER TABLE users_achievements 
        ADD FOREIGN KEY (user_id) REFERENCES users(id),
        ADD FOREIGN KEY (achievement_id) REFERENCES achievements(id);
    `
	return execQuery(db, query)
}

// Создание связей для таблицы trains
func createTrainsRelations(db *sql.DB) error {
	query := `
    ALTER TABLE trains 
        ADD FOREIGN KEY (lead_muscle_id) REFERENCES muscles(id);
    `
	return execQuery(db, query)
}

// Создание связей для таблицы exercises_help_muscle
func createExercisesHelpMuscleRelations(db *sql.DB) error {
	query := `
    ALTER TABLE exercises_help_muscle 
        ADD FOREIGN KEY (exercise_id) REFERENCES exercises(id),
        ADD FOREIGN KEY (help_muscle_id) REFERENCES muscles(id);
    `
	return execQuery(db, query)
}

// Создание связей для таблицы train_help_muscle
func createTrainHelpMuscleRelations(db *sql.DB) error {
	query := `
    ALTER TABLE train_help_muscle 
        ADD FOREIGN KEY (train_id) REFERENCES trains(id),
        ADD FOREIGN KEY (help_muscle_id) REFERENCES muscles(id);
    `
	return execQuery(db, query)
}
