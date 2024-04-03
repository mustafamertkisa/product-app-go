package model

type User struct {
	Id    uint   `gorm:"type:int;primaryKey"`
	Name  string `gorm:"not null" json:"name,omitempty"`
	Email string `gorm:"not null" json:"email,omitempty"`
}
