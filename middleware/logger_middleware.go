package middleware

import (
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggerMiddleware(app *fiber.App) {
	// membuat folder logs jika belum ada
	err := os.MkdirAll("./storage/logs", os.ModePerm)

	// cek error
	if err != nil {
		log.Fatal(err)
	}

	// membuka atau membuat file log (jika file log belum ada)
	file, err := os.OpenFile(
		"./storage/logs/app.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)

	// cek error
	if err != nil {
		log.Fatal(err)
	}

	// output terminal + file
	multiWriter := io.MultiWriter(os.Stdout, file)

	// register middleware logger
	app.Use(logger.New(logger.Config{
		Output: multiWriter,
	}))
}