package service

import (
	"mose/model/entity"
	"mose/model/response"
)

func parseGenres(titleGenres []entity.TitleGenre) (genres []response.Genre) {
	for _, titleGenre := range titleGenres {
		genre := response.Genre{
			Id:   titleGenre.GenreId,
			Name: titleGenre.Genre.Name,
		}

		genres = append(genres, genre)
	}

	return genres
}
