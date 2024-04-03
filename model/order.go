package model

type Order struct {
	Id        uint    `gorm:"type:int;primaryKey"`
	UserId    uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null" json:"quantity,omitempty"`
	User      User    `gorm:"foreignKey:UserId" json:"user,omitempty"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}
