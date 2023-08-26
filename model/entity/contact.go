package entity

type Contact struct {
	ID      int64  `json:"id"`
	Email   string `json:"email"`
	Age     int    `json:"age,omitemty"`
	Name    string `json:"nombre"`
	Address string `json:"address"`
}
