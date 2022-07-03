package domain

import "mose/model/entity"

type Movie struct {
	Title
	entity.Episode
}
