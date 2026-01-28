package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/infrastructure/config"
)

func NewDatabaseConnection(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_USER,
		cfg.DB_Password,
		cfg.DB_Name,
		cfg.DB_SSLMode,
	)
	log.Println("DSN:", dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connet database", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Connection failed to ping", err)
	}

	log.Println("Database connection successfully")
	return db
}

func NewMigrations(db *sql.DB, cfg *config.Config) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("Failed to create migration driver", err)
	}

	mig, err := migrate.NewWithDatabaseInstance(
		"file://internal/infrastructure/migrations",
		cfg.DB_Name,
		driver,
	)

	if err != nil {
		log.Fatal("Failed to create migrations", err)
	}

	if err := mig.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Failed to connect migration", err)
	}
	log.Println("Database migrations connect successfully")
}
