package main

import (
	"log"
	"mose/config"
	"mose/controller"
	"mose/repository"
	"mose/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.LoadConfig(".")
	db, err := repository.CreateConnectionPostreSql(config)

	if err != nil {
		panic("Failed to connect to database : " + err.Error())
	}

	titleRepository := repository.NewTitleRepository(db)
	accountRepository := repository.NewAccountRepository(db)

	movieService := service.NewMovieService(titleRepository)
	seriesService := service.NewSeriesService(titleRepository)
	favoriteService := service.NewFavoriteService(titleRepository)
	titleService := service.NewTitleService(titleRepository)
	accountService := service.NewAccountService(accountRepository)

	movieController := controller.NewMovieController(movieService, titleService)
	seriesController := controller.NewSeriesController(seriesService, titleService)
	favoriteController := controller.NewFavoriteController(favoriteService)
	accountController := controller.NewAccountController(accountService)
	titleController := controller.NewTitleController(titleService)

	app := fiber.New()

	movieApi := app.Group("/movies")
	movieApi.Get("/:id", movieController.GetById)
	movieApi.Get("", movieController.Search)
	movieApi.Use(accountController.VerifyAccountRole).Delete("/:id", movieController.Delete)

	seriesApi := app.Group("/series")
	seriesApi.Get("/:id", seriesController.GetById)
	seriesApi.Get("", seriesController.Search)
	seriesApi.Use(accountController.VerifyAccountRole).Delete("/:id", seriesController.Delete)

	favoriteApi := app.Group("/accounts/:accountid/favorites")
	favoriteApi.Delete("/:id", favoriteController.Delete)
	favoriteApi.Get("", favoriteController.Search)
	favoriteApi.Post("", favoriteController.Create)

	titleApi := app.Group("/titles")
	titleApi.Use(accountController.VerifyAccountRole).Post("", titleController.Create)

	log.Fatal(app.Listen(":" + config.AppPort))
}
