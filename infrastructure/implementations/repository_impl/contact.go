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

func (p *ContactRepoImpl) GetByName(name string) ([]entity.Contact, error) {
	// An contacts slice to hold data from returned rows.
	var contacts []entity.Contact
	db := p.p.AztDB.DB

	rows, err := db.Query("SELECT id,email,age,name FROM contact WHERE name = $1", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var contact entity.Contact
		if err := rows.Scan(&contact.ID, &contact.Email, &contact.Age, &contact.Name); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		contacts = append(contacts, contact)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return contacts, nil

}
func ContactRepository(p *persistence.Persistence) repository.ContactRepository {
	return &ContactRepoImpl{p}
}
