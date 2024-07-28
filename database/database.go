package database

import (
	"fmt"
	"log"
	"os"

	"github.com/illiyyin/sagala-todo/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBinstance struct {
	Db *gorm.DB
}

var DB DBinstance

func ConnectDB() {
	log.Println("Load .env file")
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	log.Println("Connecting DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode((logger.Info)),
	})

	if err != nil {
		log.Fatal("Failed to connect to Database. \n", err)
		os.Exit((2))
	}

	// command if you already migrated the database
	if err := db.AutoMigrate(&model.Task{}, &model.TaskStatus{}); err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	log.Println("Database migrated successfully!")

	log.Println("DB connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	DB = DBinstance{
		Db: db,
	}
}
