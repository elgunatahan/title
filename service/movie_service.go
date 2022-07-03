package service

import (
	"mose/errors"
	"mose/model/domain"
	"mose/model/response"
	"mose/repository"
)

type movieService struct {
	titleRepository repository.TitleRepository
}

type MovieService interface {
	GetById(id int64) (*response.Movie, errors.Error)
	Search(name string, genre uint8, page int, pagesize int) ([]response.SearchMovies, errors.Error)
}

func NewMovieService(titleRepository repository.TitleRepository) *movieService {

	return &movieService{titleRepository: titleRepository}
}

func (t movieService) GetById(id int64) (*response.Movie, errors.Error) {
	result, err := t.titleRepository.GetMovieTitleById(id)

	if err != nil {
		return nil, err
	}
	return parseMovie(result), nil
}

func parseMovie(result *domain.Movie) *response.Movie {
	stars := []response.Star{}

	for _, episodeStar := range result.Episode.EpisodeStars {
		star := response.Star{
			Id:      episodeStar.Star.Id,
			Name:    episodeStar.Star.Name,
			Surname: episodeStar.Star.Surname,
		}

		stars = append(stars, star)
	}

	movie := &response.Movie{
		Id:          result.Title.Id,
		EpisodeId:   result.Episode.Id,
		Name:        result.Title.Name,
		Type:        result.Title.Type,
		IMDBId:      result.Title.IMDBId,
		Genres:      parseGenres(result.Title.TitleGenres),
		Description: result.Title.Description,
		Rating:      result.Title.Rating,
		ReleaseDate: result.Title.ReleaseDate,
		Duration:    result.Title.Duration,
		Director:    response.Director(result.Episode.Director),
		Writer:      response.Writer(result.Episode.Writer),
		Stars:       stars,
		Audio:       result.Episode.Audio,
		Subtitles:   result.Episode.Subtitles,
	}

	return movie
}

func (t movieService) Search(name string, genre uint8, page int, pagesize int) ([]response.SearchMovies, errors.Error) {
	searchMovies := []response.SearchMovies{}

	res, err := t.titleRepository.SearchMovieTitle(name, genre, page, pagesize)
	if err != nil {
		return nil, err
	}

	for _, title := range res {
		searchMovie := response.SearchMovies{
			Id:          title.Id,
			Name:        title.Name,
			Type:        title.Type,
			IMDBId:      title.IMDBId,
			Description: title.Description,
			Rating:      title.Rating,
			ReleaseDate: title.ReleaseDate,
			Duration:    title.Duration,
			Genres:      parseGenres(title.TitleGenres),
		}

		searchMovies = append(searchMovies, searchMovie)
	}

	return searchMovies, nil
}
