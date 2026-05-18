package pinjam_buku

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// membuat object app untuk representasi dari fiber
	app := fiber.New(
		fiber.Config{
			// implementasi prefork
		// agar aplikasi bisa optimal, menyesuaikan dengan jumlah CORE CPU di server
		Prefork: true,
		},
	)

	// implementasi endpoint routing
	app.Get("/", func (ctx *fiber.Ctx) error {
		// menampilkan hello world 
		return ctx.SendString("Hello World")
	})

	// implementasi middleware
	// membuat middleware untuk melakukan loggging request (untuk semua endpoint)
	// untuk mencatat log dari request masuk dan keluar
	app.Use(func (ctx *fiber.Ctx) error {
		fmt.Println("This is middleware before processing request")
		
		// memanggil Ctx.Next, untuk melanjutkan ke proses selanjutnya
		err := ctx.Next()

		fmt.Println("This is middleware after processing request")

		return err 
	})

	// menjalankan server 
	// mengembalikan error
	err := app.Listen("localhost:3000")

	// mengecek error
	if err != nil {
		panic(err)
	}
}