package models

type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Email string `gorm:"unique;not null" json:"email"`
	Name  string `json:"name"`
}
