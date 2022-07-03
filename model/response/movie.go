package response

import "time"

type Movie struct {
	Id          int64
	EpisodeId   int64
	Name        string
	Type        uint8
	IMDBId      int64
	Genres      []Genre
	Description string
	Rating      float32
	ReleaseDate time.Time
	Duration    int64
	Director    Director
	Writer      Writer
	Stars       []Star
	Audio       string
	Subtitles   string
}
