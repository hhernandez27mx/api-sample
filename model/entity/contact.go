package entity

import "time"

type Contact struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"  `
	Email     string    `json:"email" gorm:"unique"`
	Age       int       `json:"age,omitempty" `
	Name      string    `json:"nombre" binding:"required"`
	Addresses []Address `json:"addresses"`
}

type Address struct {
	ID        int64     `json:"id"`
	Country   string    `json:"country"`
	ContactID int64     `json:"contact_id"`
	CreatedAt time.Time `json:"-"  `
	Street    string    `json:"street" `
	NumberExt *string   `json:"number_ext,omitempty"`
}
