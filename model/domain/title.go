package domain

import (
	"mose/model/entity"
	"time"
)

type Title struct {
	Id          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Type        uint8
	IMDBId      int64
	IsDeleted   bool
	Description string
	Rating      float32
	ReleaseDate time.Time
	Duration    int64
	TitleGenres []entity.TitleGenre
}
