package errors

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Error interface {
	GetStatus() int
	GetMessage() string
	Error() string
}

type errorImpl struct {
	status  int
	message string
}

func (err *errorImpl) Error() string {
	return fmt.Sprintf(err.message)
}

func (err *errorImpl) GetStatus() int {
	return err.status
}

func (err *errorImpl) GetMessage() string {
	return err.message
}

func New(status int, message string) Error {
	return &errorImpl{status: status, message: message}
}

func HandleError(ctx *fiber.Ctx, error Error) {

	ctx.JSON(error.GetStatus())

}
