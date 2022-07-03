package controller

import (
	"mose/errors"
	"mose/model/entity"
	"mose/model/request"
	"mose/service"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type titleController struct {
	titleService service.TitleService
}

func NewTitleController(titleService service.TitleService) *titleController {

	return &titleController{
		titleService: titleService,
	}
}

func (t titleController) Create(c *fiber.Ctx) error {
	request := &request.CreateTitle{}
	c.BodyParser(request)
	if err := request.Validate(); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	title := entity.Title{
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        request.Name,
		Type:        uint8(request.Type),
		IMDBId:      request.IMDBId,
		Description: request.Description,
		Rating:      request.Rating,
		ReleaseDate: request.ReleaseDate,
		Duration:    request.Duration,
		IsDeleted:   false,
	}

	err := t.titleService.Create(title)

	if err != nil {
		errors.HandleError(c, err)
		return c.Status(err.GetStatus()).JSON(err.Error())
	}

	return c.SendStatus(http.StatusOK)
}
