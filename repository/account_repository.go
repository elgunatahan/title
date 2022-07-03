package repository

import (
	"mose/errors"
	"net/http"

	"gorm.io/gorm"
)

type accountRepository struct {
	connection *gorm.DB
}

type AccountRepository interface {
	VerifyAccountRole(username string, password string, role string) errors.Error
}

func NewAccountRepository(connection *gorm.DB) *accountRepository {

	return &accountRepository{connection: connection}
}

func (t accountRepository) VerifyAccountRole(username string, password string, role string) errors.Error {
	var exists bool

	res := t.connection.Table("accounts").
		Select("count(*) > 0").
		Joins("join account_roles ar on accounts.id = ar.account_id").
		Joins("join roles r on ar.role_id = r.id").
		Where("accounts.username = ? and accounts.password = ? and r.name = ?", username, password, role).
		Find(&exists)

	if res.Error != nil {
		return errors.New(http.StatusBadRequest, "An error occured on getting account with username: "+username+", error: "+res.Error.Error())
	}

	if !exists {
		return errors.New(http.StatusNotFound, "Account with username: "+username+" unauthorized")
	}

	return nil
}
