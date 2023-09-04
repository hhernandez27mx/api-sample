package repositoryimpl

import (
	"fmt"

	"azt.com/api-sample/infrastructure/persistence"
	"azt.com/api-sample/model/entity"
	"azt.com/api-sample/model/repository"
)

type ContactRepoImpl struct {
	p *persistence.Persistence
}

func (p *ContactRepoImpl) Add(c entity.Contact) error {
	return nil
}

func (p *ContactRepoImpl) GetByName(name string) ([]entity.Contact, error) {
	var contacts []entity.Contact
	db := p.p.AztDB.DB

	err := db.Where("name = ?", name).Find(&contacts).Error
	if err != nil {
		fmt.Printf("error %q: %v", name, err)
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return contacts, nil

}
func ContactRepository(p *persistence.Persistence) repository.ContactRepository {
	return &ContactRepoImpl{p}
}
