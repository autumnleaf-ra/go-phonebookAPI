package repository

import "github.com/gofiber/fiber/v2"

func (repo *Repository) SetupRoute(app *fiber.App) {
	api := app.Group("/api/")

	/* Number */
	api.Get("/", repo.GetNumber)
	api.Post("/add", repo.AddNumber)
	api.Delete("/delete/:id", repo.DeleteNumber)
	api.Get("/detail/:id", repo.DetailNumber)

	/* Favorites */
	api2 := app.Group("/api/favorites")
	api2.Get("/", repo.GetNumberFavorites)
	api2.Post("/add/:id", repo.AddFavorites)
	api2.Delete("/delete/:id", repo.DeleteFavorites)
	api2.Get("/detail/:id", repo.DetailFavorites)
}
