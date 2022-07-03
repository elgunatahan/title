package request

import (
	"mose/util"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateTitle struct {
	Name        string
	Type        util.TitleType
	IMDBId      int64
	Description string
	Rating      float32
	ReleaseDate time.Time
	Duration    int64
}

func (r CreateTitle) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Type, validation.In(util.Movie, util.Series)),
		validation.Field(&r.IMDBId, validation.Required),
		validation.Field(&r.Description, validation.Required),
		validation.Field(&r.Rating, validation.Required),
		validation.Field(&r.ReleaseDate, validation.Required),
		validation.Field(&r.Duration, validation.Required),
	)
}
