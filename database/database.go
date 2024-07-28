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

	// you can turn off this flag if the db already migrated
	isMigrating := os.Getenv("INIT_DB")
	if isMigrating == "1" {
		migrateAndInitTaskStatus(db)
	}

	log.Println("DB connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	DB = DBinstance{
		Db: db,
	}
}

func migrateAndInitTaskStatus(db *gorm.DB) {
	log.Println("Migrating DB!")
	if err := db.AutoMigrate(&model.Task{}, &model.TaskStatus{}); err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}
	log.Println("Database migrated successfully!")

	generateTaskStatus(db)
}

func generateTaskStatus(db *gorm.DB) {
	log.Println("Init Task Status")
	taskStatuses := []model.TaskStatus{
		model.TaskStatus{
			StatusName: "Waiting List",
		},
		model.TaskStatus{
			StatusName: "In Progress",
		},
		model.TaskStatus{
			StatusName: "Done",
		},
	}
	if err := db.Create(&taskStatuses).Error; err != nil {
		fmt.Println(err)
		return
	}
	log.Println("Successfully init Task Status")
}
