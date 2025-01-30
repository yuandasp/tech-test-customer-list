package models

type Customer struct {
	ID  		 uint      `gorm:"primaryKey" json:"id"`
	Name         string    `json:"name"`
	BirthDate    string    `json:"birth_date"`
	Address      string    `json:"address"`
	PhoneNumber  string    `json:"phone_number"`
}
