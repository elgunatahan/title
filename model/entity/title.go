package entity

import "time"

type Title struct {
	Id          int64 `gorm:"primaryKey"`
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
}
