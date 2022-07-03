package service

import (
	"mose/errors"
	"mose/model/entity"
	"mose/repository"
	"mose/util"
	"net/http"
)

type titleService struct {
	titleRepository repository.TitleRepository
}

type TitleService interface {
	Create(title entity.Title) errors.Error
	Delete(id int64, titleType util.TitleType) errors.Error
}

func NewTitleService(titleRepository repository.TitleRepository) *titleService {

	return &titleService{titleRepository: titleRepository}
}

func (t titleService) Delete(id int64, titleType util.TitleType) errors.Error {

	return t.titleRepository.DeleteTitle(id, uint8(titleType))
}

func (t titleService) Create(title entity.Title) errors.Error {

	isExist, err := t.titleRepository.TitleIsExist(title.Name, title.Type)
	if err != nil {
		return err
	}

	if isExist {
		return errors.New(http.StatusConflict, "Title with name: "+title.Name+" already exist!")
	}

	return t.titleRepository.CreateTitle(title)
}
