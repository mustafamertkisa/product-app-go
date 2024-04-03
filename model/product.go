package model

type Product struct {
	Id    uint    `gorm:"type:int;primaryKey"`
	Name  string  `gorm:"not null" json:"name,omitempty"`
	Price float64 `gorm:"not null" json:"price,omitempty"`
}
