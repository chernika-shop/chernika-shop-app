package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)


func RunMigrations() error {
	dbPath := os.Getenv("DB_BASE_URL")
	
	log.Println("Connecting to the db for migrations")
	db,err := sql.Open("postgres",dbPath)

	if err != nil{
		return fmt.Errorf("failed to open db: %w",err)
	}
	defer db.Close()
	
	if err := db.Ping(); err != nil{
		return fmt.Errorf("failed to ping db: %w",err)
	}
	log.Println("All ok -> db is sigma")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	
	if err != nil {
		return fmt.Errorf("failed to create driver: %w",err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)

	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w",err)
	}

	log.Println("OKEEEEY LET'S GO RUNNING MIGRATIONS")

	if err := migration.Up(); err != nil{
		if err == migrate.ErrNoChange{
			log.Println("New migrations is ready")
			return nil
		}
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	log.Println("Migrations applied")
	return nil
}