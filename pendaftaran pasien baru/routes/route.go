package routes

import (
	"pendaftaran_pasien_baru/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteHandler(r *fiber.App) {
r.Get("/patients", handler.BiodataHandler)
r.Post("/create_patient", handler.BiodataHandlerCreate)
r.Get("/patient/:id", handler.UserHandlerGetById)
r.Put("/patient/:id", handler.UserHandlerUpdate)
r.Delete("/patient/:id", handler.UserHandlerDelete)
}