package main

import (
	"pendaftaran_pasien_baru/db"
	"pendaftaran_pasien_baru/db/migration"
	"pendaftaran_pasien_baru/routes"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Prefork:      true,
	})
	//migration
	migration.MigrationHandler()
	//database
	db.DbInit()
	//routes
	routes.RouteHandler(app)
	// if fiber.IsChild() {
	// 		fmt.Println("I'm child process")
	// 	} else {
	// 		fmt.Println("I'm parent process")
	// 	}

	err := app.Listen(":7001")
	if err != nil {
		panic(err)
	}
}