package service

import (
	"mose/errors"
	"mose/model/entity"
	"mose/model/response"
	"mose/repository"
	"net/http"
	"strconv"
)

type favoriteService struct {
	titleRepository repository.TitleRepository
}

type FavoriteService interface {
	SearchFavoriteTitle(accountId int64, name string, genre uint8, page int, pagesize int) ([]response.SearchFavorites, errors.Error)
	Delete(accountId int64, id int64) errors.Error
	Create(favorite entity.Favorite) errors.Error
}

func NewFavoriteService(titleRepository repository.TitleRepository) *favoriteService {

	return &favoriteService{titleRepository: titleRepository}
}

func (t favoriteService) SearchFavoriteTitle(accountId int64, name string, genre uint8, page int, pagesize int) ([]response.SearchFavorites, errors.Error) {
	searchFavorites := []response.SearchFavorites{}

	res, err := t.titleRepository.SearchFavoriteTitle(accountId, name, genre, page, pagesize)
	if err != nil {
		return nil, err
	}
	for _, favorite := range res {
		searchFavorite := response.SearchFavorites{
			Id:          favorite.Favorite.Id,
			TitleId:     favorite.Title.Id,
			Name:        favorite.Title.Name,
			Type:        favorite.Title.Type,
			IMDBId:      favorite.Title.IMDBId,
			Description: favorite.Title.Description,
			Rating:      favorite.Title.Rating,
			ReleaseDate: favorite.Title.ReleaseDate,
			Duration:    favorite.Title.Duration,
			Genres:      parseGenres(favorite.Title.TitleGenres),
		}

		searchFavorites = append(searchFavorites, searchFavorite)
	}
	return searchFavorites, nil
}

func (t favoriteService) Delete(accountId int64, id int64) errors.Error {

	return t.titleRepository.DeleteFavorite(accountId, id)
}

func (t favoriteService) Create(favorite entity.Favorite) errors.Error {
	isExist, err := t.titleRepository.FavoriteIsExist(favorite.TitleId, favorite.AccountId)
	if err != nil {
		return err
	}

	if isExist {
		return errors.New(http.StatusConflict, "Favorite with TitleId: "+strconv.FormatInt(favorite.TitleId, 10)+" already exist!")
	}

	return t.titleRepository.CreateFavorite(favorite)
}
