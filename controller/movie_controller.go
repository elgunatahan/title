package controller

import (
	"mose/errors"
	"mose/service"
	"mose/util"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type movieController struct {
	movieService service.MovieService
	titleService service.TitleService
}

func NewMovieController(movieService service.MovieService, titleService service.TitleService) *movieController {

	return &movieController{
		movieService: movieService,
		titleService: titleService,
	}
}

func (t movieController) GetById(c *fiber.Ctx) error {

	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	res, err := t.movieService.GetById(id)

	if err != nil {

		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(res)
}

func (t movieController) Search(c *fiber.Ctx) error {

	name := c.Query("name")
	genre, _ := strconv.ParseUint(c.Query("genre"), 10, 64)
	page, _ := strconv.Atoi(c.Query("page", "0"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize", "0"))
	res, err := t.movieService.Search(name, uint8(genre), page, pagesize)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(res)
}

func (t movieController) Delete(c *fiber.Ctx) error {

	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	err := t.titleService.Delete(id, util.Movie)

	if err != nil {
		errors.HandleError(c, err)
		return c.Status(err.GetStatus()).JSON(err.Error())
	}

	return c.SendStatus(http.StatusOK)
}
