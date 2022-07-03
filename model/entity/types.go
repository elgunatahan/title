package entity

import "time"

type Star struct {
	Id      int64 `gorm:"primaryKey"`
	Name    string
	Surname string
}

type TitleGenre struct {
	Id      int64 `gorm:"primaryKey"`
	TitleId int64
	GenreId int64
	Genre   Genre
}

type EpisodeStar struct {
	Id        int64 `gorm:"primaryKey"`
	EpisodeId int64
	StarId    int64
	Star      Star
}

type Director struct {
	Id      int64 `gorm:"primaryKey"`
	Name    string
	Surname string
}

type Writer struct {
	Id      int64 `gorm:"primaryKey"`
	Name    string
	Surname string
}

type Genre struct {
	Id   int64 `gorm:"primaryKey"`
	Name string
}

type SeasonEpisode struct {
	Id          int64 `gorm:"primaryKey"`
	TitleId     int64
	SeasonId    int64
	EpisodeId   int64
	Episode     Episode
	Name        string
	Number      int64
	Description string
	Rating      float32
	ReleaseDate time.Time
	Duration    int64
}

type Season struct {
	Id             int64 `gorm:"primaryKey"`
	TitleId        int64 `gorm:"column:title_id"`
	Number         int64
	SeasonEpisodes []SeasonEpisode
}

type Episode struct {
	Id           int64 `gorm:"primaryKey"`
	TitleId      int64
	Audio        string
	Subtitles    string
	DirectorId   int64
	Director     Director
	WriterId     int64
	Writer       Writer
	EpisodeStars []EpisodeStar
}

type AccountRole struct {
	Id        int64 `gorm:"primaryKey"`
	AccountId int64
	RoleId    int64
	Role      Role
}

type Role struct {
	Id   int64 `gorm:"primaryKey"`
	Name string
}
