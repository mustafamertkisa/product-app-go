package model

type Product struct {
	Id    int     `gorm:"primaryKey" json:"id"`
	Name  string  `gorm:"not null" json:"name,omitempty"`
	Price float64 `gorm:"not null" json:"price,omitempty"`
}
