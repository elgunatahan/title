package repository

import (
	"mose/errors"
	"mose/model/domain"
	"mose/model/entity"
	"net/http"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type titleRepository struct {
	connection *gorm.DB
}

type TitleRepository interface {
	GetMovieTitleById(id int64) (*domain.Movie, errors.Error)
	SearchMovieTitle(name string, genre uint8, page int, pagesize int) ([]domain.Title, errors.Error)
	GetSeriesTitleById(id int64) ([]domain.Series, errors.Error)
	SearchSeriesTitle(name string, genre uint8, page int, pagesize int) ([]domain.Title, errors.Error)
	SearchFavoriteTitle(accountId int64, name string, genre uint8, page int, pagesize int) ([]domain.Favorite, errors.Error)
	DeleteTitle(id int64, titleType uint8) errors.Error
	DeleteFavorite(accountId int64, favoriteId int64) errors.Error
	CreateFavorite(favorite entity.Favorite) errors.Error
	CreateTitle(title entity.Title) errors.Error
	TitleIsExist(name string, titleType uint8) (bool, errors.Error)
	FavoriteIsExist(titleId int64, accountId int64) (bool, errors.Error)
}

func NewTitleRepository(connection *gorm.DB) *titleRepository {

	return &titleRepository{connection: connection}
}

func (t titleRepository) GetMovieTitleById(id int64) (*domain.Movie, errors.Error) {
	response := &domain.Movie{}
	res := t.connection.Table("titles").
		Joins("JOIN episodes on titles.id=episodes.title_id").
		Where("titles.Type = 0 and is_deleted = false").
		Select("titles.*,episodes.*").
		Preload("TitleGenres").
		Preload("TitleGenres.Genre").
		Preload("Director").
		Preload("Writer").
		Preload("EpisodeStars").
		Preload("EpisodeStars.Star").
		Find(&response, "titles.id = ?", id)

	if res.Error != nil {
		return nil, errors.New(http.StatusBadRequest, "An error occured on getting movie with Id: "+strconv.FormatInt(id, 10)+", error:"+res.Error.Error())
	}

	if res.RowsAffected == 0 {
		return nil, errors.New(http.StatusNotFound, "Movie with id: "+strconv.FormatInt(id, 10)+" not found")
	}

	return response, nil
}

func (t titleRepository) SearchMovieTitle(name string, genre uint8, page int, pagesize int) ([]domain.Title, errors.Error) {
	response := []domain.Title{}
	query := t.connection.Table("titles").Joins("JOIN title_genres on titles.id=title_genres.title_id").Where("Type = 0 and is_deleted = false")

	if len(name) > 0 {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%")
	}

	if genre > 0 {
		query = query.Where("title_genres.genre_id = ?", genre)
	}

	query = query.Select("distinct titles.*").Preload("TitleGenres").Preload("TitleGenres.Genre")

	if page > 0 && pagesize > 0 {
		query.Offset((page - 1) * pagesize).Limit(pagesize)
	}

	res := query.Find(&response)

	if res.Error != nil {
		return nil, errors.New(http.StatusBadRequest, "An error occured on searching movie, error: "+res.Error.Error())
	}

	return response, nil
}

func (t titleRepository) GetSeriesTitleById(id int64) ([]domain.Series, errors.Error) {
	response := []domain.Series{}

	query := t.connection.Table("titles")

	query = query.
		Joins("JOIN seasons on titles.id=seasons.title_id").
		Where("Type = 1 and is_deleted = false")

	query = query.Select("titles.*, seasons.*")

	res := query.
		Preload("TitleGenres").
		Preload("TitleGenres.Genre").
		Preload("SeasonEpisodes").
		Preload("SeasonEpisodes.Episode").
		Preload("SeasonEpisodes.Episode.Director").
		Preload("SeasonEpisodes.Episode.Writer").
		Preload("SeasonEpisodes.Episode.EpisodeStars").
		Preload("SeasonEpisodes.Episode.EpisodeStars.Star").
		Find(&response, "titles.id = ?", id)

	if res.Error != nil {
		return nil, errors.New(http.StatusBadRequest, "An error occured on getting series with Id: "+strconv.FormatInt(id, 10)+", error:"+res.Error.Error())
	}

	if res.RowsAffected == 0 {
		return nil, errors.New(http.StatusNotFound, "Series with id: "+strconv.FormatInt(id, 10)+" not found")
	}

	return response, nil
}

func (t titleRepository) SearchSeriesTitle(name string, genre uint8, page int, pagesize int) ([]domain.Title, errors.Error) {
	response := []domain.Title{}
	query := t.connection.Table("titles").
		Joins("JOIN title_genres on titles.id=title_genres.title_id").
		Where("Type = 1 and is_deleted = false")

	if len(name) > 0 {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%")
	}

	if genre > 0 {
		query = query.Where("title_genres.genre_id = ?", genre)
	}

	query = query.Select("distinct titles.*").Preload("TitleGenres").Preload("TitleGenres.Genre")

	if page > 0 && pagesize > 0 {
		query.Offset((page - 1) * pagesize).Limit(pagesize)
	}

	res := query.Find(&response)

	if res.Error != nil {
		return nil, errors.New(http.StatusBadRequest, "An error occured on searching series, error: "+res.Error.Error())
	}

	if res.RowsAffected == 0 {
		return []domain.Title{}, nil
	}

	return response, nil
}

func (t titleRepository) SearchFavoriteTitle(accountId int64, name string, genre uint8, page int, pagesize int) ([]domain.Favorite, errors.Error) {
	response := []domain.Favorite{}
	query := t.connection.Table("favorites").
		Joins("JOIN titles on favorites.title_id=titles.id").
		Joins("JOIN title_genres on titles.id=title_genres.title_id").
		Where("favorites.account_id = ? and titles.is_deleted = false", accountId)

	if len(name) > 0 {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%")
	}

	if genre > 0 {
		query = query.Where("title_genres.genre_id = ?", genre)
	}

	query = query.Select("distinct favorites.*, titles.*").Preload("TitleGenres").Preload("TitleGenres.Genre")

	if page > 0 && pagesize > 0 {
		query.Offset((page - 1) * pagesize).Limit(pagesize)
	}

	res := query.Find(&response)

	if res.Error != nil {
		return nil, errors.New(http.StatusBadRequest, "An error occured on searching favorites, error: "+res.Error.Error())
	}

	if res.RowsAffected == 0 {
		return []domain.Favorite{}, nil
	}

	return response, nil
}

func (t titleRepository) DeleteTitle(id int64, titleType uint8) errors.Error {
	res := t.connection.
		Table("titles").
		Where("titles.id = ? and titles.type = ? and is_deleted = ?", id, titleType, false).
		Updates(map[string]interface{}{"is_deleted": true, "updated_at": time.Now().UTC()})

	if res.Error != nil {
		return errors.New(http.StatusBadRequest, "An error occured on deleting title with Id: "+strconv.FormatInt(id, 10)+", error: "+res.Error.Error())
	}

	if res.RowsAffected == 0 {
		return errors.New(http.StatusNotFound, "Title with id: "+strconv.FormatInt(id, 10)+" not found")
	}

	return nil
}

func (t titleRepository) DeleteFavorite(accountId int64, favoriteId int64) errors.Error {
	res := t.connection.
		Table("favorites").
		Where("favorites.id = ? and favorites.account_id = ?", favoriteId, accountId).
		Delete(&entity.Favorite{})

	if res.Error != nil {
		return errors.New(http.StatusBadRequest, "An error occured on deleting favorite with Id: "+strconv.FormatInt(favoriteId, 10)+", error: "+res.Error.Error())
	}

	if res.RowsAffected == 0 {
		return errors.New(http.StatusNotFound, "Favorite with id: "+strconv.FormatInt(favoriteId, 10)+" not found")
	}

	return nil
}

func (t titleRepository) CreateFavorite(favorite entity.Favorite) errors.Error {
	res := t.connection.Create(&favorite)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New(http.StatusInternalServerError, "An error occured on creating favorite with TitleId: "+strconv.FormatInt(favorite.TitleId, 10)+", error: "+res.Error.Error())
	}

	return nil
}

func (t titleRepository) CreateTitle(title entity.Title) errors.Error {
	res := t.connection.Create(&title)
	if res.Error != nil || res.RowsAffected == 0 {
		return errors.New(http.StatusInternalServerError, "An error occured on creating tilte with name: "+title.Name+", error: "+res.Error.Error())
	}

	return nil
}

func (t titleRepository) TitleIsExist(name string, titleType uint8) (bool, errors.Error) {
	res := t.connection.
		Table("titles").
		Where("titles.name = ? and titles.type = ?", name, titleType).
		Find(&entity.Title{})

	if res.Error != nil {
		return false, errors.New(http.StatusBadRequest, "An error occured on getting title with name: "+name+", error: "+res.Error.Error())
	}

	return res.RowsAffected != 0, nil
}

func (t titleRepository) FavoriteIsExist(titleId int64, accountId int64) (bool, errors.Error) {

	res := t.connection.
		Table("favorites").
		Where("favorites.title_id = ? and favorites.account_id = ?", titleId, accountId).
		Find(&entity.Favorite{})

	if res.Error != nil {
		return false, errors.New(http.StatusBadRequest, "An error occured on getting favorite with TitleId: "+strconv.FormatInt(titleId, 10)+", error: "+res.Error.Error())
	}

	return res.RowsAffected != 0, nil
}
