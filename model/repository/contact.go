package repository

import "azt.com/api-sample/model/entity"

type ContactRepository interface {
	GetByName(string) ([]entity.Contact, error)
	Add(c entity.Contact) error
}
