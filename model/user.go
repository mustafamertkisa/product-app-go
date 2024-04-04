package model

type User struct {
	Id    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"not null" json:"name,omitempty"`
	Email string `gorm:"not null" json:"email,omitempty"`
}
