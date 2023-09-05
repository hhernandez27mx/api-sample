package services

import (
	repositoryimpl "azt.com/api-sample/infrastructure/implementations/repository_impl"
	"azt.com/api-sample/infrastructure/persistence"
	"azt.com/api-sample/model/entity"
)

type ContactService struct {
	Persistence *persistence.Persistence
}

func (c *ContactService) AddServ(contact entity.Contact) error {
	contactRepo := repositoryimpl.ContactRepository(c.Persistence)
	err := contactRepo.Add(contact)

	return err
}

func (c *ContactService) GetByName(name string) []entity.Contact {
	contactRepo := repositoryimpl.ContactRepository(c.Persistence)
	data, err := contactRepo.GetByName(name)
	if err != nil {
		return []entity.Contact{}
	}
	return data
}

func NewContactService(p *persistence.Persistence) *ContactService {
	return &ContactService{p}
}
