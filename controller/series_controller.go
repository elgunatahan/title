package controller

import (
	"mose/errors"
	"mose/service"
	"mose/util"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type seriesController struct {
	seriesService service.SeriesService
	titleService  service.TitleService
}

func NewSeriesController(seriesService service.SeriesService, titleService service.TitleService) *seriesController {

	return &seriesController{
		seriesService: seriesService,
		titleService:  titleService,
	}
}

func (t seriesController) GetById(c *fiber.Ctx) error {

	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	res, err := t.seriesService.GetSeriesTitleById(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(res)

}

func (t seriesController) Search(c *fiber.Ctx) error {
	name := c.Query("name")
	genre, _ := strconv.ParseUint(c.Query("genre"), 10, 64)
	page, _ := strconv.Atoi(c.Query("page", "0"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize", "0"))
	res, err := t.seriesService.SearchSeriesTitle(name, uint8(genre), page, pagesize)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(res)

}

func (t seriesController) Delete(c *fiber.Ctx) error {

	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	err := t.titleService.Delete(id, util.Series)

	if err != nil {
		errors.HandleError(c, err)
		return c.Status(err.GetStatus()).JSON(err.Error())
	}

	return c.SendStatus(http.StatusOK)
}
