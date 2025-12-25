package database

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunMigrations() error {
	dbURL, migrationsDir, err := prepareMigrations()
	if err != nil {
		return err
	}

	if err := ensureInitialMigration(dbURL, migrationsDir); err != nil {
		return err
	}

	return applyMigrations(dbURL)
}

func prepareMigrations() (string, string, error) {
	dbURL := os.Getenv("DB_BASE_URL")
	if dbURL == "" {
		return "", "", fmt.Errorf("DB_BASE_URL environment variable is not set")
	}

	log.Println("Connecting to the database for migrations")

	migrationsDir, err := filepath.Abs("migrations")
	if err != nil {
		return "", "", fmt.Errorf("failed to get migrations directory path: %w", err)
	}

	// Создаем директорию миграций/проверяем есть ли она
	if err := os.MkdirAll(migrationsDir, 0755); err != nil {
		return "", "", fmt.Errorf("failed to create migrations directory: %w", err)
	}

	return dbURL, migrationsDir, nil
}

func ensureInitialMigration(dbURL, migrationsDir string) error {
	// Проверяем наличие существующих миграций
	files, err := filepath.Glob(filepath.Join(migrationsDir, "*.sql"))
	if err != nil {
		return fmt.Errorf("failed to check migrations: %w", err)
	}

	// Если миграций нет, создаем начальную из GORM моделей
	if len(files) == 0 {
		return createInitialMigration(dbURL)
	}

	return nil
}

func createInitialMigration(dbURL string) error {
	log.Println("No migrations found. Creating initial migration from models")

	cmd := exec.Command("atlas", "migrate", "diff",
		"initial",
		"--env", "dev",
		"--dir", "file://migrations",
	)
	cmd.Dir = "."
	cmd.Env = append(os.Environ(), fmt.Sprintf("DB_BASE_URL=%s", dbURL))

	output, err := cmd.CombinedOutput()
	if err != nil {
		// Если ошибка из-за нечистой БД (уже есть schema_migrations), пропускаем создание
		if strings.Contains(string(output), "not clean") {
			log.Println("Database already has schema. Skipping initial migration creation.")
			log.Println("If you need to create migrations, clean the database first or create them manually.")
			return nil // Пропускаем создание, так как БД уже инициализирована
		}

		log.Printf("Atlas migrate diff output: %s", string(output))
		return fmt.Errorf("failed to create initial migration: %w (output: %s)", err, string(output))
	}

	log.Printf("Initial migration created: %s", string(output))
	return nil
}

func applyMigrations(dbURL string) error {
	log.Println("Applying migrations with Atlas")

	cmd := exec.Command("atlas", "migrate", "apply",
		"--dir", "file://migrations",
		"--url", dbURL,
		"--baseline", "", // Разрешаем применение на нечистой БД
	)
	cmd.Dir = "."

	output, err := cmd.CombinedOutput()
	if err != nil {
		// Если ошибка из-за нечистой БД и миграций нет, это нормально
		if strings.Contains(string(output), "not clean") && strings.Contains(string(output), "baseline") {
			log.Println("Database is not clean and no migrations to apply. This is normal for existing databases.")
			return nil
		}

		log.Printf("Atlas migrate apply output: %s", string(output))
		return fmt.Errorf("failed to apply migrations: %w (output: %s)", err, string(output))
	}

	log.Println("Migrations applied successfully")
	return nil
}
