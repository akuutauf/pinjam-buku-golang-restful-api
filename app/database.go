package app

import (
	"fmt"
	"os"
	"pinjam-buku/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// membuat koneksi baru untuk ke database
// db di set sebagai pointer karena berjenis struct
// untuk data source name dari parameter open bisa mereferensikan dari dokumentasi berikut ini :-
// https://github.com/go-sql-driver/mysql

func NewDB() *gorm.DB {
	// load env
	err := godotenv.Load()

	// mengecek error load file env
	helper.PanicIfError(err)

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
		// mengaktifkan log bawaan
		Logger: logger.Default.LogMode(logger.Info),

		// performance: skip default transaction dan cache statement
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	// mengecek error
	helper.PanicIfError(err)

	// ambil native sql.DB
	sqlDB, err := db.DB()

	// connection pooling
	sqlDB.SetMaxIdleConns(5)                   // maksimal koneksi ketika iddle adalah 5
	sqlDB.SetMaxOpenConns(20)                  // maksimal koneksi yang terbuka pertama kali adalah 20
	sqlDB.SetConnMaxLifetime(60 * time.Minute) // maksimal waktu koneksi berjalan adalah 60 menit
	sqlDB.SetConnMaxIdleTime(10 * time.Minute) // maksimal koneksi ketika iddle adalah 10 menit

	return db
}
