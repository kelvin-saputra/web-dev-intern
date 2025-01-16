package config

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	postgresDriver "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB //Using Class Attribute to make these variables (DB Connection) accessible to all functions in the package

func InitPsqlDB() {
	err := godotenv.Load(".env.local")

	if err != nil {
		log.Fatalf("Error while loading .env file: %s", err)
	}

	dataSource := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"), os.Getenv("DB_TIMEZONE"))

	databaseConn, err := gorm.Open(postgres.Open(dataSource), &gorm.Config{})

	DB = databaseConn
	if err != nil {
		log.Fatalf("Error while connecting to database: %s", err)
	}
}

func MigratePsqlDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error while getting database connection: %s", err)
	}
	sqlDB.Exec("DEALLOCATE ALL")

	driver, err := postgresDriver.WithInstance(sqlDB, &postgresDriver.Config{})
	if err != nil {
		log.Fatalf("Error while creating database driver instance: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)

	if err != nil {
		log.Fatalf("Error while creating migrate instance: %s", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error while running migrations: %s", err)
	}
}
