package service

import (
	"mose/errors"
	"mose/model/domain"
	"mose/model/response"
	"mose/repository"
)

type seriesService struct {
	titleRepository repository.TitleRepository
}

type SeriesService interface {
	GetSeriesTitleById(id int64) (*response.Series, errors.Error)
	SearchSeriesTitle(name string, genre uint8, page int, pagesize int) ([]response.SearchSeries, errors.Error)
}

func NewSeriesService(titleRepository repository.TitleRepository) *seriesService {

	return &seriesService{titleRepository: titleRepository}
}

func (t seriesService) GetSeriesTitleById(id int64) (*response.Series, errors.Error) {
	result, err := t.titleRepository.GetSeriesTitleById(id)
	if err != nil {
		return nil, err
	}
	series := parseSeries(result)
	return series, nil
}

func parseSeries(result []domain.Series) (series *response.Series) {
	title := result[0].Title
	genres := parseGenres(title.TitleGenres)
	seasons := []response.Season{}
	directorMap := make(map[int]response.Director)
	writerMap := make(map[int]response.Writer)
	starMap := make(map[int]response.Star)

	for _, series := range result {
		season := response.Season{
			Id:       series.Season.Id,
			Number:   series.Season.Number,
			Episodes: []response.Episode{},
		}

		for _, seasonEpisode := range series.SeasonEpisodes {
			episode := response.Episode{
				Id:          seasonEpisode.Episode.Id,
				Name:        seasonEpisode.Name,
				Number:      seasonEpisode.Number,
				Description: seasonEpisode.Description,
				Rating:      seasonEpisode.Rating,
				ReleaseDate: seasonEpisode.ReleaseDate,
				Duration:    seasonEpisode.Duration,
				Audio:       seasonEpisode.Episode.Audio,
				Subtitles:   seasonEpisode.Episode.Subtitles,
				Director:    response.Director(seasonEpisode.Episode.Director),
				Writer:      response.Writer(seasonEpisode.Episode.Writer),
			}
			directorMap[int(episode.Director.Id)] = episode.Director
			writerMap[int(episode.Writer.Id)] = episode.Writer

			stars := []response.Star{}
			for _, episodeStar := range seasonEpisode.Episode.EpisodeStars {
				star := response.Star{
					Id:      episodeStar.Star.Id,
					Name:    episodeStar.Star.Name,
					Surname: episodeStar.Star.Surname,
				}

				stars = append(stars, star)
				starMap[int(star.Id)] = star
			}
			episode.Stars = stars

			season.Episodes = append(season.Episodes, episode)
		}

		seasons = append(seasons, season)
	}

	series = &response.Series{
		Id:          title.Id,
		Name:        title.Name,
		Type:        title.Type,
		IMDBId:      title.IMDBId,
		Genres:      genres,
		Description: title.Description,
		Rating:      title.Rating,
		ReleaseDate: title.ReleaseDate,
		Duration:    title.Duration,
		Seasons:     seasons,
		Directors:   directorMap,
		Writers:     writerMap,
		Stars:       starMap,
	}

	return series
}

func (t seriesService) SearchSeriesTitle(name string, genre uint8, page int, pagesize int) ([]response.SearchSeries, errors.Error) {
	result := []response.SearchSeries{}
	res, err := t.titleRepository.SearchSeriesTitle(name, genre, page, pagesize)
	if err != nil {
		return nil, err
	}
	for _, title := range res {
		searchSeries := response.SearchSeries{
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

		result = append(result, searchSeries)
	}
	return result, nil
}
