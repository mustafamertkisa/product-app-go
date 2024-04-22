package model

type User struct {
	Id    int    `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"not null" json:"name,omitempty"`
	Email string `gorm:"not null" json:"email,omitempty"`
}
