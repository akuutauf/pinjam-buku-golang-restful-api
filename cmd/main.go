package main

import (
	"pinjam-buku/internal/database"
	"pinjam-buku/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// koneksi database
	db := database.OpenConnection()

	// hapus semua tabel (jika diperlukan)
	// database.DropAllTables(db)

	// running auto migration
	database.RunMigration(db)
	
	// membuat object app untuk representasi dari fiber
	app := fiber.New(
		fiber.Config{
			// implementasi prefork
		// agar aplikasi bisa optimal, menyesuaikan dengan jumlah CORE CPU di server
		Prefork: true,
		},
	)

	// implementasi middleware logger
	middlewares.LoggerMiddleware(app)

	// implementasi endpoint routing
	app.Get("/", func (ctx *fiber.Ctx) error {
		// menampilkan hello world 
		return ctx.SendString("Hello World")
	})

	// menjalankan server 
	// mengembalikan error
	err := app.Listen("localhost:3000")

	// mengecek error
	if err != nil {
		panic(err)
	}
}