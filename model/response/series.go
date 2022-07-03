package response

import "time"

type Series struct {
	Id          int64
	Name        string
	Type        uint8
	IMDBId      int64
	Genres      []Genre
	Description string
	Rating      float32
	ReleaseDate time.Time
	Duration    int64
	Directors   map[int]Director
	Writers     map[int]Writer
	Stars       map[int]Star
	Seasons     []Season
}

type Genre struct {
	Id   int64
	Name string
}

type Season struct {
	Id       int64
	Number   int64
	Episodes []Episode
}

type Episode struct {
	Id          int64
	Name        string
	Number      int64
	Description string
	Rating      float32
	ReleaseDate time.Time
	Duration    int64
	Audio       string
	Subtitles   string
	Director    Director
	Writer      Writer
	Stars       []Star
}
