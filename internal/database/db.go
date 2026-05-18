package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	// load env
	err := godotenv.Load()

	// mengecek error load file env
	if err != nil {
		log.Fatal("Failed to load file .env")
	}

	// ambil env
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	
	// membuat destinasi untuk database yang dituju
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	// open connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// logger query sql
		Logger: logger.Default.LogMode(logger.Info),

		// performance: skip default transaction dan cache statement
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	// cek error
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// ambil native sql db
	sqlDB, err := db.DB()

	// cek error
	if err != nil {
		log.Fatalf("Failed to get sql.DB instance: %v", err)
	}

	// connection pool
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	log.Println("Successfully connected to the database")

	return db
}