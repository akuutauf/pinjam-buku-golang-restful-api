package middleware

import (
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggerMiddleware(app *fiber.App) {

	// create logs directory
	err := os.MkdirAll("storage/logs", os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	// open/create log file
	file, err := os.OpenFile(
		"storage/logs/app.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)

	if err != nil {
		log.Fatal(err)
	}

	// test write
	_, err = file.WriteString("Logger initialized...\n")

	if err != nil {
		log.Fatal(err)
	}

	// sync file
	file.Sync()

	// multi output
	multiWriter := io.MultiWriter(os.Stdout, file)

	// register middleware
	app.Use(logger.New(logger.Config{
		// konfigurasi log request dari API
		Output:     multiWriter,
		Format:     "[${time}] ${status} | ${latency} | ${method} | ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))
}
