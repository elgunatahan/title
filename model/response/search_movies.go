package response

import "time"

type SearchMovies struct {
	Id          int64
	Name        string
	Type        uint8
	IMDBId      int64
	Genres      []Genre
	Description string
	Rating      float32
	ReleaseDate time.Time
	Duration    int64
}
