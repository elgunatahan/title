package service

import (
	"mose/errors"
	"mose/repository"
)

type accountService struct {
	accountRepository repository.AccountRepository
}

type AccountService interface {
	VerifyAccountRole(username string, password string, role string) errors.Error
}

func NewAccountService(accountRepository repository.AccountRepository) *accountService {

	return &accountService{accountRepository: accountRepository}
}

func (t accountService) VerifyAccountRole(username string, password string, role string) errors.Error {

	err := t.accountRepository.VerifyAccountRole(username, password, role)
	if err != nil {
		return err
	}

	return nil
}
