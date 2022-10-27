package routes

import (
	"github.com/baguseka01/golang_fiber_blog/controllers"
	"github.com/baguseka01/golang_fiber_blog/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsAuthenticate)
	app.Post("/api/post", controllers.CreatePost)
	app.Get("/api/allpost", controllers.AllPost)
	app.Get("/api/detailpost/:id", controllers.DetailPost)
	app.Put("/api/updatepost/:id", controllers.UpdatePost)
	app.Get("/api/uniquepost", controllers.UniquePost)
	app.Delete("/api/uniquepost/:id", controllers.DeletePost)
	app.Post("/api/uploadimage", controllers.UploadImage)
	app.Static("/api/uploads", "./uploads")
}
