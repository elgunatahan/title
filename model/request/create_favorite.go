package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateFavorite struct {
	TitleId int64
}

func (r CreateFavorite) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.TitleId, validation.Required, validation.Min(1)),
	)
}
