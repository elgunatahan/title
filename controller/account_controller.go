package controller

import (
	"encoding/base64"
	"mose/service"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type accountController struct {
	accountService service.AccountService
}

func NewAccountController(accountService service.AccountService) *accountController {

	return &accountController{
		accountService: accountService,
	}
}

func (t accountController) VerifyAccountRole(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).SendString("Please provide valid auth credentials")
	}
	token := strings.Split(authHeader, "Basic ")
	if len(token) != 2 {
		return c.Status(401).SendString("Please provide valid auth credentials")
	}

	decodedByteValue, err := base64.StdEncoding.DecodeString(token[1])
	if err != nil {
		return c.Status(401).SendString("Please provide valid auth credentials")
	}
	decodedValue := string(decodedByteValue)
	if decodedValue == "" {
		return c.Status(401).SendString("Please provide valid auth credentials")
	}

	credentials := strings.Split(decodedValue, ":")
	if len(credentials) != 2 || credentials[0] == "" || credentials[1] == "" {
		return c.Status(401).SendString("Please provide valid auth credentials")
	}

	err = t.accountService.VerifyAccountRole(credentials[0], credentials[1], "ContentManager")

	if err != nil {
		return c.Status(401).JSON(err.Error())
	}

	return c.Next()
}
