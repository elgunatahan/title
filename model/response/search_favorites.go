package response

import "time"

type SearchFavorites struct {
	Id          int64
	TitleId     int64
	Name        string
	Type        uint8
	IMDBId      int64
	Genres      []Genre
	Description string
	Rating      float32
	ReleaseDate time.Time
	Duration    int64
}
