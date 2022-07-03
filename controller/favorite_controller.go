package controller

import (
	"mose/errors"
	"mose/model/entity"
	"mose/model/request"
	"mose/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type favoriteController struct {
	favoriteService service.FavoriteService
}

func NewFavoriteController(favoriteService service.FavoriteService) *favoriteController {

	return &favoriteController{
		favoriteService: favoriteService,
	}
}

func (t favoriteController) Search(c *fiber.Ctx) error {
	accountId, _ := strconv.ParseInt(c.Params("accountid"), 10, 64)
	name := c.Query("name")
	genre, _ := strconv.ParseUint(c.Query("genre"), 10, 64)
	page, _ := strconv.Atoi(c.Query("page", "0"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize", "0"))

	res, err := t.favoriteService.SearchFavoriteTitle(accountId, name, uint8(genre), page, pagesize)

	if err != nil {
		errors.HandleError(c, err)
		return c.Status(err.GetStatus()).JSON(err.Error())
	}

	return c.JSON(res)

}

func (t favoriteController) Delete(c *fiber.Ctx) error {

	accountId, _ := strconv.ParseInt(c.Params("accountid"), 10, 64)
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	err := t.favoriteService.Delete(accountId, id)

	if err != nil {
		errors.HandleError(c, err)
		return c.Status(err.GetStatus()).JSON(err.Error())
	}

	return c.SendStatus(http.StatusOK)
}

func (t favoriteController) Create(c *fiber.Ctx) error {

	accountId, error := strconv.ParseInt(c.Params("accountid"), 10, 64)
	if error != nil {
		c.Status(400).SendString("Please provide accountid")
	}

	request := &request.CreateFavorite{}
	c.BodyParser(request)
	if err := request.Validate(); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	favorite := entity.Favorite{
		TitleId:   request.TitleId,
		AccountId: accountId,
	}

	err := t.favoriteService.Create(favorite)

	if err != nil {
		errors.HandleError(c, err)
		return c.Status(err.GetStatus()).JSON(err.Error())
	}

	return c.SendStatus(http.StatusOK)
}
